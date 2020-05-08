package declaration_test

import (
	"encoding/json"
	"github.com/vusalalishov/rapit/core/declaration"
	"github.com/vusalalishov/rapit/core/http"
	"github.com/vusalalishov/rapit/core/model"
	"github.com/vusalalishov/rapit/test/utils"
	"testing"
)

func TestEnrichSimpleCase(t *testing.T) {

	response := prepareResponse(` { "data": { "a": "b", "c": "d" } } `, t)
	decl := prepareDeclaration(` { "a": "f", "c": "g" } `, t)
	extract := prepareExtractBlock(` { "a": "{{.data.c}}", "c": "{{.data.a}}", "b": "{{.data.a}}" } `, t)

	err := declaration.Enrich(response, decl, extract)
	failOnError(err, t)

	if (*decl)["a"] != "d" || (*decl)["b"] != "b" || (*decl)["c"] != "b" {
		t.Fatal(*decl)
	}

}

func TestEnrichDeepExtraction(t *testing.T) {
	response := prepareResponse(` { "data": { "a": "b", "c": "d" } } `, t)
	decl := prepareDeclaration(` 
		{ 
			"a": {
				"b": {
					"c": {
						"d" : {
							"k": "l",
							"f": "e"
						}
					}
				}
			}
		} 
	`, t)
	extract := prepareExtractBlock(` 
		{ 
			"a": {
				"b": {
					"c": {
						"d" : {
							"k": "{{.data.a}}",
							"f": "{{.data.c}}",
							"g": "h"
						}
					}
				}
			}
		} 
	`, t)

	err := declaration.Enrich(response, decl, extract)
	failOnError(err, t)

	kVal, err := utils.NestedMapLookup(*decl, "a", "b", "c", "d", "k")
	failOnError(err, t)

	fVal, err := utils.NestedMapLookup(*decl, "a", "b", "c", "d", "f")
	failOnError(err, t)

	gVal, err := utils.NestedMapLookup(*decl, "a", "b", "c", "d", "g")
	failOnError(err, t)

	if kVal != "b" || fVal != "d" || gVal != "h" {
		t.Fail()
	}

}

func prepareExtractBlock(s string, t *testing.T) *model.Extract {
	var e model.Extract
	err := json.Unmarshal([]byte(s), &e)
	failOnError(err, t)
	return &e
}

func prepareDeclaration(s string, t *testing.T) *model.Declaration {
	var decl model.Declaration
	err := json.Unmarshal([]byte(s), &decl)
	failOnError(err, t)
	return &decl
}

func prepareResponse(body string, t *testing.T) *http.Response {
	var responseBody interface{}
	err := json.Unmarshal([]byte(body), &responseBody)

	failOnError(err, t)

	return &http.Response{
		Body: &responseBody,
	}
}

func failOnError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
