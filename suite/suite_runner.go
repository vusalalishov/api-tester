package suite

import (
	"api-tester/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func RunSuite(suite model.Suite) {

	declarationMap := suite.Declare

	for _, testCase := range suite.Cases {
		runCase(testCase, declarationMap)
	}
}

func runCase(testCase model.Case, declaration model.Declaration) {

	for _, scenario := range testCase.Scenarios {
		var try = scenario.Try
		response, err := sendRequest(try, declaration)

		if err != nil {
			panic(err)
		}

		var verifyScenario = scenario.Verify
		message, err := verifyResponse(response, verifyScenario)

		if err != nil {
			fmt.Println(message)
			break
		}
	}

}

func verifyResponse(response *HttpResponse, scenario model.VerifyScenario) (string, error) {
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

func sendRequest(scenario model.TryScenario, declaration model.Declaration) (*HttpResponse, error) {
	payload, err := json.Marshal(scenario.Payload)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(payload)

	resp, err := http.Post(scenario.Url, "application/json", reader)

	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseBody map[string]interface{}

	err = json.Unmarshal(bodyBytes, &responseBody)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		body:   &responseBody,
		status: model.HttpStatus(resp.StatusCode),
	}, nil

}
