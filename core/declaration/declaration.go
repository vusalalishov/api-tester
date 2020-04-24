package declaration

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/vusalalishov/api-tester/core/http"
	"github.com/vusalalishov/api-tester/core/model"
	"reflect"
	"text/template"
)

func Enrich(response *http.Response, declaration *model.Declaration, extract *model.Extract) (err error) {
	extractBytes, err := json.Marshal(extract)

	if err != nil {
		return err
	}

	tmpl, err := template.New("extractTemplate").Parse(string(extractBytes))

	if err != nil {
		return err
	}

	var newDeclarationBytes bytes.Buffer

	if err = tmpl.Execute(&newDeclarationBytes, response.Body); err != nil {
		return err
	}

	var newDeclaration model.Declaration

	if err = json.Unmarshal(newDeclarationBytes.Bytes(), &newDeclaration); err != nil {
		return err
	}

	if err = mergeNewDeclarationsIntoExisting(declaration, &newDeclaration); err != nil {
		return err
	}

	return nil

}

func mergeNewDeclarationsIntoExisting(existing *model.Declaration, additional *model.Declaration) (err error) {

	if additional == nil {
		return nil
	}

	existingMap := map[string]interface{}(*existing)
	additionalMap := map[string]interface{}(*additional)

	for additionalKey, additionalValue := range additionalMap {
		valueKind := reflect.TypeOf(additionalValue).Kind()

		if valueKind == reflect.Map {
			existingValue, found := existingMap[additionalKey]
			if found {
				// update
				n := model.Declaration(additionalValue.(map[string]interface{}))
				e := model.Declaration(existingValue.(map[string]interface{}))
				if err = mergeNewDeclarationsIntoExisting(&e, &n); err != nil {
					return err
				}
			} else {
				// add
				existingMap[additionalKey] = additionalValue
			}
		} else if valueKind == reflect.String {
			existingMap[additionalKey] = additionalValue
		} else {
			return errors.New("additionalValue is neither type or string")
		}

	}

	return nil

}
