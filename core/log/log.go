package log

import (
	"fmt"
	"github.com/vusalalishov/api-tester/core/http"
	"github.com/vusalalishov/api-tester/core/model"
)

func RunningSuite(title string) {
	fPrintln("- Running suite: %s ...", title)
}

func RunningTestCase(title string) {
	fPrintln("-- Running test case: %s ...", title)
}

func RunningScenario(title string) {
	fPrintln("--- Running scenario: %s ...", title)
}

func ScenarioFailed(scenario model.Scenario, response *http.Response, reason error) {
	fPrintln("--- Scenario: %s is FAILED with reason: [%s]", scenario.Title, reason.Error())
}

func CaseCompleted(testCase *model.Case, isPassed bool, reason error) {
	if isPassed {
		fPrintln("-- Test case: %s is PASSED", testCase.Title)
	} else {
		fPrintln("-- Test case: %s is FAILED with reason: [%s]", testCase.Title, reason.Error())
	}
}

func ScenarioPassed(scenario model.Scenario) {
	fPrintln("--- Scenario: %s is PASSED", scenario.Title)
}

func SuiteCompleted(suite model.Suite, isPassed bool) {
	if isPassed {
		fPrintln("- Suite: %s is PASSED", suite.Title)
	} else {
		fPrintln("- Suite: %s is FAILED", suite.Title)
	}
}

func fPrintln(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}
