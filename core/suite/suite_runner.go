package suite

import (
	"errors"
	"fmt"
	"github.com/vusalalishov/api-tester/core/http"
	"github.com/vusalalishov/api-tester/core/log"
	"github.com/vusalalishov/api-tester/core/model"
	"github.com/vusalalishov/api-tester/core/verify"
)

func RunSuite(suite model.Suite, templateDir string, scriptDir string) *log.TestSuite {
	suiteLog := log.NewSuite(suite.Title)
	for _, testCase := range suite.Cases {
		suiteLog.AddMessage(fmt.Sprintf("Running test case %s", testCase.Title))
		testCaseLog := suiteLog.AddCase(testCase.Title)
		err := runCase(&testCase, &suite.Declaration, templateDir, scriptDir, testCaseLog)
		if err != nil {
			suiteLog.SetStatus(log.FAILED)
		}
	}
	return suiteLog
}

func runCase(testCase *model.Case, declaration *model.Declaration, templateDir string, scriptDir string, caseLog *log.TestCase) (err error) {

	for _, scenario := range testCase.Scenarios {

		scenarioLog := caseLog.AddScenario(scenario.Title)

		try := scenario.Try
		response, err := sendRequest(&try, declaration, templateDir)

		if err != nil {
			scenarioLog.SetStatus(log.FAILED)
			break
		}

		enrichDeclaration(response, declaration)
		errorArr := verifyResponse(response, &scenario.Verify, scriptDir, declaration)

		if len(errorArr) != 0 {
			scenarioLog.SetStatus(log.FAILED)
			break
		}
	}

	return err
}

func sendRequest(scenario *model.TryScenario, declaration *model.Declaration, templateDir string) (*http.Response, error) {
	request, err := prepareHttpRequest(scenario, declaration, templateDir)
	if err != nil {
		return nil, err
	}
	return request.Execute()
}

func enrichDeclaration(response *http.Response, declaration *model.Declaration) {
	// TODO: implementation is missing
}

func verifyResponse(response *http.Response, scenario *model.VerifyScenario, scriptDir string, declaration *model.Declaration) []error {
	if response.Status != http.Status(scenario.Status) {
		return []error{errors.New("status mismatch")}
	}
	return verify.Schema(response.Body, scenario.Schema, scriptDir, make([]error, 0))
}
