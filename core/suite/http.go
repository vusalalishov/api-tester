package suite

import (
	"bytes"
	"encoding/json"
	"github.com/vusalalishov/api-tester/core/http"
	"github.com/vusalalishov/api-tester/core/model"
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
		payload, err := json.Marshal(scenario.Payload)

		if err != nil {
			return nil, err
		}

		reader = bytes.NewReader(payload)
	}

	// TODO: add headers
	return &http.Request{
		Url:    scenario.Url,
		Method: httpMethod,
		Body:   reader,
	}, nil
}
