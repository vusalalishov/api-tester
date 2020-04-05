package model

type Protocol string
type HttpMethod string
type HttpStatus string

type Declaration map[string]interface{}

type TryScenario struct {
	Method  HttpMethod
	Url     string
	Payload interface{}
}

type VerifyScenario struct {
	Status HttpStatus
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
	Title    string      `json:"title"`
	Label    []string    `json:"Label"`
	Protocol Protocol    `json:"protocol"`
	Declare  Declaration `json:"declare"`
	Cases    []Case      `json:"cases"`
}
