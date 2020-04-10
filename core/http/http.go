package http

import (
	"bytes"
	"errors"
)

type Method string

const (
	POST Method = "POST"
	GET  Method = "GET"
)

type Header struct {
	Key   string
	Value string
}

type Request struct {
	Headers []Header
	Method  Method
	Body    *bytes.Reader
}

type Response struct {
	Status int
	Body   *map[string]interface{}
}

func (r *Request) Execute() (*Response, error) {

	return nil, errors.New("unimplemented")
}

/*

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
*/
