package suite

import (
	"bytes"
	"github.com/vusalalishov/api-tester/core/config"
	"github.com/vusalalishov/api-tester/core/http"
	"github.com/vusalalishov/api-tester/core/model"
	"io/ioutil"
	"text/template"
)

func prepareHttpRequest(scenario *model.TryScenario, declaration *model.Declaration) (*http.Request, error) {
	method := scenario.Method

	var httpMethod http.Method
	if method == model.POST {
		httpMethod = http.POST
	} else if method == model.GET {
		httpMethod = http.GET
	}

	var reader *bytes.Reader
	if scenario.Payload != nil {
		var err error
		reader, err = preparePayloadReader(scenario, declaration)

		if err != nil {
			return nil, err
		}
	}

	// TODO: add headers : https://github.com/vusalalishov/api-tester/issues/10
	return &http.Request{
		Url:    scenario.Url,
		Method: httpMethod,
		Body:   reader,
	}, nil
}

func preparePayloadReader(scenario *model.TryScenario, declaration *model.Declaration) (*bytes.Reader, error) {

	fileBytes, err := ioutil.ReadFile(config.PayloadDir(*scenario.Payload))

	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("payload").Parse(string(fileBytes))

	if err != nil {
		return nil, err
	}

	var payloadBytes bytes.Buffer

	err = tmpl.Execute(&payloadBytes, declaration)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(payloadBytes.Bytes()), nil
}
