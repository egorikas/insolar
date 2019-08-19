//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package apiclientdeprecated

import (
	"encoding/json"
	"log"
)

//Status
type NodeStatus struct {
	Result nodeStatusResult `json:"result"`
}

type nodeStatusResult struct {
	NetworkState    string     `json:"NetworkState"`
	Origin          nodeInfo   `json:"Origin"`
	ActiveListSize  int        `json:"ActiveListSize"`
	WorkingListSize int        `json:"WorkingListSize"`
	Nodes           []nodeInfo `json:"Nodes"`
	PulseNumber     int        `json:"PulseNumber"`
	Entropy         string     `json:"Entropy"`
	NodeState       string     `json:"NodeState"`
	Version         string     `json:"Version"`
}

type nodeInfo struct {
	Reference string `json:"Reference"`
	Role      string `json:"Role"`
	IsWorking bool   `json:"IsWorking"`
}

func GetNodeStatus(url string) NodeStatus {
	// Form a request body for getStatus:
	getStatusReq := PlatformRequest{
		JSONRPC: JSONRPCVersion,
		Method:  "node.getStatus",
		ID:      id,
	}
	// Increment the id for future requests:
	id++

	// Create a variable to unmarshal the seed response into:
	var status NodeStatus
	// Send the seed request:   payload, url, "POST", "application/json"
	seedRespBody := SendPlatformRequest(getStatusReq, url, "POST", "application/json")
	// Unmarshal the response:
	err := json.Unmarshal(seedRespBody, &status)
	if err != nil {
		log.Fatalln(err)
	}
	// Put the current seed into a variable:
	return status
}
