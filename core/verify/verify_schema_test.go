package verify_test

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"testing"
)

type Response struct {
	Message string
}

func TestSchema(t *testing.T) {
	vm := otto.New()

	_, err := vm.ToValue(map[string]interface{}{
		"errorCode":    "0",
		"errorMessage": "",
		"data": map[string]interface{}{
			"userId": "1",
		},
	})
	if err != nil {
		fmt.Println(err)
	}

}
