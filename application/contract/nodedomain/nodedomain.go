/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package nodedomain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/insolar/insolar/application/proxy/noderecord"
	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/cryptohelpers/ecdsa"
	"github.com/insolar/insolar/logicrunner/goplugin/foundation"
)

// NodeDomain holds noderecords
type NodeDomain struct {
	foundation.BaseContract
}

// NewNodeDomain create new NodeDomain
func NewNodeDomain() (*NodeDomain, error) {
	return &NodeDomain{}, nil
}

func (nd *NodeDomain) getNodeRecord(ref core.RecordRef) *noderecord.NodeRecord {
	return noderecord.GetObject(ref)
}

func (nd *NodeDomain) makeCertificate(numberOfBootstrapNodes int, publicKey string, majorityRule int, roles []string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["majority_rule"] = majorityRule
	result["public_key"] = publicKey
	result["roles"] = roles

	bNodes, err := nd.makeBootstrapNodesConfig(numberOfBootstrapNodes)
	if err != nil {
		return nil, fmt.Errorf("Can't make bootstrap nodes config: %s", err.Error())
	}

	result["bootstrap_nodes"] = bNodes

	return result, nil
}

func (nd *NodeDomain) makeBootstrapNodesConfig(numberOfBootstrapNodes int) ([]map[string]string, error) {

	if numberOfBootstrapNodes == 0 {
		return []map[string]string{}, nil
	}

	nodeRefs, err := nd.GetChildrenTyped(noderecord.GetClass())
	if err != nil {
		return nil, &foundation.Error{S: "[ makeBootstrapNodesConfig ] Problem with taking records: " + err.Error()}
	}

	requiredNodesNum := numberOfBootstrapNodes

	var result []map[string]string
	for _, ref := range nodeRefs {
		if requiredNodesNum == 0 {
			break
		}
		requiredNodesNum -= 1

		nodeRecord := noderecord.GetObject(ref)
		recordInfo, err := nodeRecord.GetNodeInfo()
		if err != nil {
			return nil, &foundation.Error{S: "[ makeBootstrapNodesConfig ] Can't get NodeInfo"}
		}

		bConf := map[string]string{}
		bConf["public_key"] = recordInfo.PublicKey
		bConf["host"] = recordInfo.IP

		result = append(result, bConf)
	}

	if requiredNodesNum != 0 {
		return nil, &foundation.Error{S: "[ makeBootstrapNodesConfig ] There no enough nodes"}
	}

	return result, nil
}

// RegisterNode registers node in system
func (nd *NodeDomain) RegisterNode(publicKey string, numberOfBootstrapNodes int, majorityRule int, roles []string, ip string) ([]byte, error) {
	const majorityPercentage = 0.51

	if majorityRule != 0 {
		if float32(majorityRule) <= majorityPercentage*float32(numberOfBootstrapNodes) {
			return nil, errors.New("majorityRule must be more than 0.51 * numberOfBootstrapNodes")
		}
	}

	result, err := nd.makeCertificate(numberOfBootstrapNodes, publicKey, majorityRule, roles)
	if err != nil {
		return nil, errors.New("[ RegisterNode ] " + err.Error())
	}

	// TODO: what should be done when record already exists?
	newRecord := noderecord.NewNodeRecord(publicKey, roles, ip)
	record := newRecord.AsChild(nd.GetReference())

	result["reference"] = record.GetReference().String()

	rawCert, err := json.Marshal(result)
	if err != nil {
		return nil, errors.New("Can't marshal certificate: " + err.Error())
	}

	return rawCert, nil
}

// RemoveNode deletes node from registry
func (nd *NodeDomain) RemoveNode(nodeRef core.RecordRef) error {
	node := nd.getNodeRecord(nodeRef)
	return node.Destroy()
}

// IsAuthorized checks is signature correct
func (nd *NodeDomain) IsAuthorized(nodeRef core.RecordRef, seed []byte, signatureRaw []byte) (bool, error) {
	pubKey, err := nd.getNodeRecord(nodeRef).GetPublicKey()
	if err != nil {
		return false, err
	}
	ok, err := ecdsa.Verify(seed, signatureRaw, pubKey)
	if err != nil {
		return false, err
	}
	return ok, nil
}

// Authorize checks node and returns node info
func (nd *NodeDomain) Authorize(nodeRef core.RecordRef, seed []byte, signatureRaw []byte) (pubKey string, roles []core.NodeRole, errS error) {
	// TODO: this should be removed when proxies stop panic
	defer func() {
		if r := recover(); r != nil {
			pubKey = ""
			roles = nil
			err, ok := r.(error)
			errTxt := ""
			if ok {
				errTxt = err.Error()
			}
			err = errors.New("[ Authorize ] Recover after panic: " + errTxt)
		}
	}()
	nodeR := nd.getNodeRecord(nodeRef)
	nodeInfo, err := nodeR.GetNodeInfo()
	if err != nil {
		return "", nil, fmt.Errorf("[ Authorize ] Problem with Getting info: " + err.Error())
	}

	pubKey = nodeInfo.PublicKey
	roles = nodeInfo.Roles

	ok, err := ecdsa.Verify(seed, signatureRaw, pubKey)
	if err != nil {
		return "", nil, fmt.Errorf("[ Authorize ] Problem with verifying of signature: " + err.Error())
	}
	if !ok {
		return "", nil, errors.New("[ Authorize ] Can't verify signature")
	}

	return pubKey, roles, nil
}
