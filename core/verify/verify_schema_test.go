package verify_test

import (
	"encoding/json"
	"github.com/vusalalishov/api-tester/core/verify"
	"testing"
)

func TestSchema_WillRecognizeSimpleMatch(t *testing.T) {
	testSuccessPath([]byte(`{ "a": "b" }`), t)
}

func TestSchema_WillRecognizeComplexMatch(t *testing.T) {
	bytes := []byte(
		`{ 
			"a": "b",
			"c": {
				"k": {
					"g": {
						"d": "f",
						"b": "c"
					}
				}
			}
		}`)

	testSuccessPath(bytes, t)
}

func TestSchema_WillVerifyLists(t *testing.T) {
	bytes := []byte(
		`[
			{
				"a": "b",
				"c": [
					{
						"a": "k"
					},
					"b"
				]
			},
			"c",
			{
				"d": "f",
				"b": "j"
			},
			"a",
			"b"
		]`)

	testSuccessPath(bytes, t)
}

func TestSchema_WillFindComplexMismatch(t *testing.T) {
	schemaBytes := []byte(
		`[
			{
				"a": "b",
				"c": [
					{
						"a": 1
					},
					"b"
				]
			},
			"c",
			{
				"d": "f",
				"b": "j"
			},
			"a",
			"b"
		]`)

	responseBytes := []byte(
		`[
			{
				"a": "b",
				"c": [
					{
						"a": "1"
					},
					"b"
				]
			},
			"c",
			{
				"d": "f",
				"b0": "j"
			},
			"a",
			"b"
		]`)

	testUnmatchedSchemaPath(schemaBytes, responseBytes, t)
}

func testSuccessPath(bytes []byte, t *testing.T) {

	var (
		response interface{}
		schema   interface{}
	)

	unmarshal(bytes, &response, t)
	unmarshal(bytes, &schema, t)

	err := verify.Schema(&response, &schema, []error{})

	if len(err) != 0 {
		t.Fail()
	}
}

func testUnmatchedSchemaPath(schemaBytes []byte, responseBytes []byte, t *testing.T) {

	var (
		response interface{}
		schema   interface{}
	)

	unmarshal(schemaBytes, &response, t)
	unmarshal(responseBytes, &schema, t)

	err := verify.Schema(&response, &schema, []error{})

	if len(err) == 0 {
		t.Fail()
	}
}

func unmarshal(bytes []byte, res interface{}, t *testing.T) {
	err := json.Unmarshal(bytes, res)
	if err != nil {
		t.Fatal(err, "unable to unmarshal json")
	}
}
