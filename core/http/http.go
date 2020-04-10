package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Method string
type Status int

const (
	POST Method = "POST"
	GET  Method = "GET"
)

type Header struct {
	Key   string
	Value string
}

type Request struct {
	Url     string
	Headers []Header
	Method  Method
	Body    *bytes.Reader
}

type Response struct {
	Status Status
	Body   *map[string]interface{}
}

func (r *Request) Execute() (response *Response, err error) {

	var httpResponse *http.Response

	if r.Method == POST {
		httpResponse, err = http.Post(r.Url, "application/json", r.Body)
		if err != nil {
			return nil, err
		}
	} else if r.Method == GET {
		httpResponse, err = http.Get(r.Url)
		if err != nil {
			return nil, err
		}
	}

	if httpResponse != nil {

		bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return nil, err
		}

		var responseBody map[string]interface{}

		err = json.Unmarshal(bodyBytes, &responseBody)
		if err != nil {
			return nil, err
		}

		return &Response{
			Body:   &responseBody,
			Status: Status(httpResponse.StatusCode),
		}, nil

	} else {
		return nil, errors.New("http.Response is nil")
	}

}
