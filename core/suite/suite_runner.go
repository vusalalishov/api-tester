package suite

import (
	"errors"
	"fmt"
	decl "github.com/vusalalishov/api-tester/core/declaration"
	"github.com/vusalalishov/api-tester/core/http"
	"github.com/vusalalishov/api-tester/core/log"
	"github.com/vusalalishov/api-tester/core/model"
	"github.com/vusalalishov/api-tester/core/verify"
)

func RunSuite(suite model.Suite) *log.TestSuite {
	suiteLog := log.NewSuite(suite.Title)
	for _, testCase := range suite.Cases {
		suiteLog.AddMessage(fmt.Sprintf("Running test case %s", testCase.Title))
		testCaseLog := suiteLog.AddCase(testCase.Title)
		if suite.Declaration == nil {
			d := model.Declaration(make(map[string]interface{}))
			suite.Declaration = &d
		}
		err := runCase(&testCase, suite.Declaration, "payloads/", "scripts/", testCaseLog)
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

		enrichErr := decl.Enrich(response, declaration, scenario.Extract)
		if enrichErr != nil {
			scenarioLog.SetStatus(log.FAILED)
			break
		}

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

func verifyResponse(response *http.Response, scenario *model.VerifyScenario, scriptDir string, declaration *model.Declaration) []error {
	if response.Status != http.Status(scenario.Status) {
		return []error{errors.New("status mismatch")}
	}
	return verify.Schema(response.Body, scenario.Schema, scriptDir, make([]error, 0))
}
