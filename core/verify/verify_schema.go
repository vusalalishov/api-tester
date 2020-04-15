package verify

import (
	"errors"
	"fmt"
	"reflect"
)

// TODO: refactor it
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
		responseMap, responseIsStringMap := (*response).(map[string]interface{})
		schemaMap, schemaIsStringMap := (*schema).(map[string]interface{})
		if !responseIsStringMap || !schemaIsStringMap {
			failures = append(failures, errors.New("response or schema is not string map"))
		}
		failures = append(failures, verifyMap(&responseMap, &schemaMap, failures)...)
	} else if responseKind == reflect.Array || responseKind == reflect.Slice {
		responseArray, isResponsePresent := (*response).([]interface{})
		schemaArray, isSchemaPresent := (*schema).([]interface{})
		if isResponsePresent && isSchemaPresent {
			failures = append(failures, verifyArray(&responseArray, &schemaArray, failures)...)
		} else {
			failures = append(failures, errors.New("schema mismatch"))
		}

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
		responseValue, ok := (*responseMap)[key]
		if ok {
			failures = append(failures, Schema(&responseValue, &schemaValue, failures)...)
		} else {
			failures = append(failures, errors.New("not found by key: "+key))
		}
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
