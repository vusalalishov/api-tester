package suite

import "api-tester/model"

type HttpResponse struct {
	status model.HttpStatus
	body   *map[string]interface{}
}
