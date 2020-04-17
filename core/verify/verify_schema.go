package verify

import (
	"errors"
	"github.com/robertkrimen/otto"
	"github.com/vusalalishov/api-tester/core/model"
	"os"
)

func Schema(response *interface{}, schema model.Schema, scriptDir string, failures []error) []error {
	// read js file
	if schema.Tests != nil {
		for file, testMethod := range *schema.Tests {
			script, err := os.Open(scriptDir + file)
			if err == nil {
				vm := otto.New()
				_, err := vm.Run(script)

				if err != nil {
					failures = append(failures, err)
					continue
				}

				fn, err := vm.Get(testMethod)

				if err != nil {
					failures = append(failures, err)
					continue
				}

				val, err := vm.ToValue(*response)

				if err != nil {
					failures = append(failures, err)
					continue
				}

				testResult, err := fn.Call(otto.NullValue(), val)

				if err != nil {
					failures = append(failures, err)
					continue
				}

				exitCodeVal, err := testResult.Object().Get("exitCode")

				if err != nil {
					failures = append(failures, err)
					continue
				}

				exitCode, err := exitCodeVal.ToInteger()

				if err != nil {
					failures = append(failures, err)
					continue
				}

				if exitCode != 0 {
					failures = append(failures, errors.New("exit code is not 0"))
				}

			} else {
				failures = append(failures, err)
			}
		}
	}
	return failures
}
