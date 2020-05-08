package suite

import (
	"errors"
	decl "github.com/vusalalishov/rapit/core/declaration"
	"github.com/vusalalishov/rapit/core/http"
	"github.com/vusalalishov/rapit/core/log"
	"github.com/vusalalishov/rapit/core/model"
	"github.com/vusalalishov/rapit/core/verify"
)

func RunSuite(suite model.Suite) {
	log.RunningSuite(suite.Title)
	suiteIsPassed := true
	for _, testCase := range suite.Cases {
		log.RunningTestCase(testCase.Title)
		if suite.Declaration == nil {
			d := model.Declaration(make(map[string]interface{}))
			suite.Declaration = &d
		}
		caseIsPassed, err := runCase(&testCase, suite.Declaration)
		log.CaseCompleted(&testCase, caseIsPassed, err)
		if !caseIsPassed {
			suiteIsPassed = false
		}
	}
	log.SuiteCompleted(suite, suiteIsPassed)
}

func runCase(testCase *model.Case, declaration *model.Declaration) (bool, error) {

	var isFailed = false
	var err error = nil

	for _, scenario := range testCase.Scenarios {

		log.RunningScenario(scenario.Title)

		try := scenario.Try
		response, sendingError := sendRequest(&try, declaration)

		if sendingError != nil {
			log.ScenarioFailed(scenario, response, sendingError)
			err = sendingError
			isFailed = true
			break
		}

		enrichErr := decl.Enrich(response, declaration, scenario.Extract)
		if enrichErr != nil {
			log.ScenarioFailed(scenario, response, enrichErr)
			isFailed = true
			err = enrichErr
			break
		}

		errorArr := verifyResponse(response, &scenario.Verify, declaration)

		if len(errorArr) != 0 {
			isFailed = true
			err = errors.New("response verification error")
			log.ScenarioFailed(scenario, response, err)
			break
		}

		log.ScenarioPassed(scenario)
	}

	return !isFailed, err
}

func sendRequest(scenario *model.TryScenario, declaration *model.Declaration) (*http.Response, error) {
	request, err := prepareHttpRequest(scenario, declaration)
	if err != nil {
		return nil, err
	}
	return request.Execute()
}

func verifyResponse(response *http.Response, scenario *model.VerifyScenario, declaration *model.Declaration) []error {
	if response.Status != http.Status(scenario.Status) {
		return []error{errors.New("status mismatch")}
	}
	return verify.Schema(response.Body, scenario.Schema, make([]error, 0))
}
