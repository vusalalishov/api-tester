package model

type Protocol string
type HttpMethod string
type HttpStatus int

type Declaration map[string]interface{}

type HttpHeaders map[string]interface{}

type TryScenario struct {
	Method  HttpMethod
	Headers HttpHeaders
	Url     string
	Payload interface{}
}

type VerifyScenario struct {
	Status HttpStatus `json:",string"`
	Schema *map[string]interface{}
}

type Scenario struct {
	Try    TryScenario
	Verify VerifyScenario
}

type Case struct {
	Title     string
	Labels    []string
	Scenarios []Scenario
}

type Suite struct {
	Title       string
	Label       []string
	Protocol    Protocol
	Declaration Declaration `json:"declare"`
	Cases       []Case
}
