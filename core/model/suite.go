package model

type Protocol string
type HttpMethod string
type HttpStatus int

type Declaration map[string]interface{}

type Extract map[string]interface{}

type HttpHeaders map[string]interface{}
type Tests map[string]string

type TryScenario struct {
	Method      HttpMethod
	Headers     HttpHeaders
	Url         string
	Payload     *string
	PayloadFile *string
}

type Schema struct {
	Tests *Tests
}

type VerifyScenario struct {
	Status  HttpStatus `json:",string"`
	Headers HttpHeaders
	Schema  Schema
}

type Scenario struct {
	Title   string
	Try     TryScenario
	Extract *Extract
	Verify  VerifyScenario
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
	Declaration *Declaration `json:"declare"`
	Cases       []Case
}
