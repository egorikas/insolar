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
