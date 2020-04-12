package verify

import (
	"errors"
	"fmt"
	"reflect"
)

func Schema(response *interface{}, schema *interface{}, failures []error) []error {
	if response == nil && schema == nil {
		return failures
	}

	if response == nil || schema == nil {
		failures = append(failures, errors.New("response or schema are null"))
		return failures
	}

	responseKind := reflect.TypeOf(*response).Kind()
	schemaKind := reflect.TypeOf(*schema).Kind()

	if responseKind != schemaKind {
		failures = append(failures, errors.New("response and schema type mismatch"))
		return failures
	}

	if responseKind == reflect.Map {
		responseMap := (*response).(map[string]interface{})
		schemaMap := (*schema).(map[string]interface{})
		failures = append(failures, verifyMap(&responseMap, &schemaMap, failures)...)
	} else if responseKind == reflect.Array || responseKind == reflect.Slice {
		responseArray := (*response).([]interface{})
		schemaArray := (*schema).([]interface{})
		failures = append(failures, verifyArray(&responseArray, &schemaArray, failures)...)
	} else {
		// perform equality check
		if *response != *schema {
			failures = append(failures, errors.New(fmt.Sprintf("mismatch: response = %b, schema = %b", response, schema)))
		}
	}
	return failures
}

func verifyMap(responseMap *map[string]interface{}, schemaMap *map[string]interface{}, failures []error) []error {
	for key, schemaValue := range *schemaMap {
		responseValue := (*responseMap)[key]
		failures = append(failures, Schema(&responseValue, &schemaValue, failures)...)
	}
	return failures
}

func verifyArray(responseArray *[]interface{}, schemaArray *[]interface{}, failures []error) []error {
	if len(*schemaArray) != len(*responseArray) {
		failures = append(failures, errors.New("array size mismatch"))
		return failures
	}
	for index, schema := range *schemaArray {
		response := (*responseArray)[index]
		failures = append(failures, Schema(&response, &schema, failures)...)
	}
	return failures
}
