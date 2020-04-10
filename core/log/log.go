package log

import "github.com/vusalalishov/api-tester/core/http"

type Status int

const (
	SUCCESS Status = 0
	FAILED  Status = 1
)

type testInfo struct {
	Status   Status
	Messages []*string
}

type TestSuite struct {
	Title string
	Cases []*TestCase
	testInfo
}

type TestCase struct {
	Title     string
	Scenarios []*TestScenario
	testInfo
}

type TestScenario struct {
	Title string
	testInfo
	OriginalPayload http.Response
}
