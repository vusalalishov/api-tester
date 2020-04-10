package suite

import (
	"api-tester/http"
	"api-tester/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

func RunSuite(suite model.Suite) {
	for _, testCase := range suite.Cases {
		runCase(&testCase, &suite.Declaration)
	}
}

func runCase(testCase *model.Case, declaration *model.Declaration) {

	for _, scenario := range testCase.Scenarios {
		try := scenario.Try
		response, err := sendRequest(&try, declaration)

		if err != nil {
			panic(err)
		}

		verifyScenario := scenario.Verify
		message, err := verifyResponse(response, verifyScenario)

		if err != nil {
			fmt.Println(message)
			break
		}
	}

}

func sendRequest(scenario *model.TryScenario, declaration *model.Declaration) (*http.Response, error) {
	request := prepareHttpRequest(scenario, declaration)
	response, err := request.Execute()
	return response, err
}

func verifyResponse(response *http.Response, scenario model.VerifyScenario) (string, error) {
	body := response.body
	return verify(*body, *scenario.Schema)
}

func verify(body map[string]interface{}, schema map[string]interface{}) (string, error) {

	for key, value := range schema {

		v := body[key]

		if reflect.TypeOf(v).Kind() == reflect.Map {
			s, err := verify(v.(map[string]interface{}), value.(map[string]interface{}))
			if err != nil {
				return s, err
			}
		} else {
			if v.(string) != value.(string) {
				return "Verification failed", errors.New("failed")
			}
		}
	}

	return "", nil
}
