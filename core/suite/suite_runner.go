package suite

import (
	"api-tester/core/http"
	"api-tester/core/log"
	"api-tester/core/model"
	"errors"
	"fmt"
	"reflect"
)

func RunSuite(suite model.Suite) *log.TestSuite {
	suiteLog := log.NewSuite(suite.Title)
	for _, testCase := range suite.Cases {
		suiteLog.AddMessage(fmt.Sprintf("Running test case %s", testCase.Title))
		testCaseLog := suiteLog.AddCase(testCase.Title)
		err := runCase(&testCase, &suite.Declaration, testCaseLog)
		if err != nil {
			suiteLog.SetStatus(log.FAILED)
		}
	}
	return suiteLog
}

func runCase(testCase *model.Case, declaration *model.Declaration, caseLog *log.TestCase) (err error) {

	for _, scenario := range testCase.Scenarios {

		scenarioLog := caseLog.AddScenario(scenario.Title)

		try := scenario.Try
		response, err := sendRequest(&try, declaration)

		if err != nil {
			scenarioLog.SetStatus(log.FAILED)
			break
		}

		enrichDeclaration(response, declaration)
		err = verifyResponse(response, &scenario.Verify, declaration)

		if err != nil {
			break
		}
	}

	return err
}

func sendRequest(scenario *model.TryScenario, declaration *model.Declaration) (*http.Response, error) {
	request := prepareHttpRequest(scenario, declaration)
	response, err := request.Execute()
	return response, err
}

func enrichDeclaration(response *http.Response, declaration *model.Declaration) {
	// TODO: implementation is missing
}

func verifyResponse(response *http.Response, scenario *model.VerifyScenario, declaration *model.Declaration) error {
	if response.Status != int(scenario.Status) {
		return errors.New("status mismatch")
	}
	return verify(response.Body, scenario.Schema)
}

func verify(body *map[string]interface{}, schema *map[string]interface{}) error {

	for key, value := range *schema {

		v := (*body)[key]

		if reflect.TypeOf(v).Kind() == reflect.Map {
			err := verify(v.(*map[string]interface{}), value.(*map[string]interface{}))
			if err != nil {
				return err
			}
		} else {
			if v.(string) != value.(string) {
				return errors.New("field mismatch")
			}
		}
	}

	return nil
}
