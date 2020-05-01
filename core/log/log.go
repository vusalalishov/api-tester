package log

import (
	"fmt"
	"github.com/vusalalishov/api-tester/core/http"
)

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

func (s *TestSuite) Print() {
	fmt.Println(s.Title)
	fmt.Println(s.Status)

	for _, testCase := range s.Cases {
		fmt.Println(testCase.Title)
		fmt.Println(testCase.Status)

		for _, scenario := range testCase.Scenarios {
			fmt.Println(scenario.Title)
			fmt.Println(scenario.Status)
			for _, message := range scenario.Messages {
				fmt.Println(message)
			}
		}

	}
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
