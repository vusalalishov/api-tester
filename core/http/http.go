package http

import (
	"bytes"
	"encoding/json"
	"github.com/ddo/rq"
	"github.com/ddo/rq/client"
	"net/http"
	"strings"
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
	Status  Status
	Body    *interface{}
	Headers []Header
}

func (r *Request) Execute() (response *Response, err error) {

	var httpResponse *http.Response

	httpRequest := rq.New(string(r.Method), r.Url)
	httpRequest.Set("Content-Type", "application/json")

	if r.Headers != nil {
		for _, h := range r.Headers {
			httpRequest.Set(h.Key, h.Value)
		}
	}

	if r.Body != nil {
		httpRequest.SendRaw(r.Body)
	}

	responseBytes, httpResponse, err := client.Send(httpRequest, true)
	if err != nil {
		return nil, err
	}

	var responseBody interface{}

	err = json.Unmarshal(responseBytes, &responseBody)
	if err != nil {
		return nil, err
	}

	responseHeaders := make([]Header, 0)

	for k, v := range httpResponse.Header {
		responseHeaders = append(responseHeaders, Header{
			Key:   k,
			Value: strings.Join(v, ""),
		})
	}

	return &Response{
		Body:   &responseBody,
		Status: Status(httpResponse.StatusCode),
	}, nil

}
