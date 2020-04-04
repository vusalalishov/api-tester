package model

type Protocol int
type HttpMethod int
type HttpStatus int

type Declaration interface{}

type TryScenario struct {
	method  HttpMethod
	url     string
	payload interface{}
}

type VerifyScenario struct {
	status HttpStatus
	schema interface{}
}

type Scenario struct {
	try    TryScenario
	verify VerifyScenario
}

type Case struct {
	title     string
	labels    []string
	scenarios []Scenario
}

type Suite struct {
	title    string
	label    []string
	protocol Protocol
	declare  Declaration
	cases    []Case
}
