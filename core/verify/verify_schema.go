package verify

import (
	"errors"
	"github.com/robertkrimen/otto"
	"github.com/vusalalishov/rapit/core/config"
	"github.com/vusalalishov/rapit/core/http"
	"github.com/vusalalishov/rapit/core/model"
	"os"
)

// TODO: this works for now, let's move on. Will get back to it when will have time for refactoring, see https://github.com/vusalalishov/rapit/issues/8
func Schema(response *http.Response, declaration *model.Declaration, schema model.Schema, failures []error) []error {
	// read js file
	if schema.Tests != nil {
		for file, testMethod := range *schema.Tests {
			script, err := os.Open(config.ScriptDir(file))
			if err == nil {
				// create JS VM
				vm := otto.New()

				// run the script - so the functions are created
				if err := runScript(vm, script, failures); err != nil {
					continue
				}

				// get the test method
				fn, err := vm.Get(testMethod)

				if err != nil {
					failures = append(failures, err)
					continue
				}

				// turn response into JS object
				jsResponseVal, err := vm.ToValue(*response)

				if err != nil {
					failures = append(failures, err)
					continue
				}

				// turn response into JS object
				jsDeclarationVal, err := vm.ToValue(*declaration)

				if err != nil {
					failures = append(failures, err)
					continue
				}

				// call the test method on response
				testResult, err := fn.Call(otto.NullValue(), jsDeclarationVal, jsResponseVal)

				if err != nil {
					failures = append(failures, err)
					continue
				}

				// get the exitCodeValue from response
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
				break
			}
		}
	}
	return failures
}

func runScript(vm *otto.Otto, script *os.File, failures []error) error {
	_, err := vm.Run(script)

	if err != nil {
		failures = append(failures, err)
		return err
	}
	return nil
}
