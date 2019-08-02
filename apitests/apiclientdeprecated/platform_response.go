package apiclientdeprecated

// Declare structures to unmarshal the responses into in accordance with the specification.
// The Platform uses the basic JSON RPC 2.0 response structure:
type PlatformResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

// The result of the seed request is as follows:
type seedResponse struct {
	PlatformResponse
	Result seedResult `json:"result"`
}

type seedResult struct {
	Seed    string `json:"Seed"`
	TraceID string `json:"TraceID"`
}

// The result of the info requests is as follows:
type infoResponse struct {
	PlatformResponse
	Result infoResult `json:"result"`
}

type infoResult struct {
	RootDomain             string   `json:"RootDomain"`
	RootMember             string   `json:"RootMember"`
	MigrationAdminMember   string   `json:"MigrationAdminMember"`
	MigrationDaemonMembers []string `json:"MigrationDaemonMembers"`
	NodeDomain             string   `json:"NodeDomain"`
	TraceID                string   `json:"TraceID"`
}

type ContractAnswer struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  *Result `json:"result,omitempty"`
	Error   *Error  `json:"error,omitempty"`
}

type Result struct {
	ContractResult interface{} `json:"callResult,omitempty"`
	TraceID        string      `json:"traceID,omitempty"`
}

type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    Data   `json:"data,omitempty"`
}

type Data struct {
	TraceID string `json:"traceID,omitempty"`
}
