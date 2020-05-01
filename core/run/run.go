package run

import (
	"encoding/json"
	"github.com/vusalalishov/api-tester/core/config"
	"github.com/vusalalishov/api-tester/core/model"
	"github.com/vusalalishov/api-tester/core/suite"
	"io/ioutil"
	"os"
	"strings"
)

func AllSuites() error {
	fileInfos, err := ioutil.ReadDir(config.GetSuiteDir())

	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		if isJsonFile(fileInfo) {
			err := Suite(fileInfo.Name())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Suite(suiteName string) error {

	suitePath := config.SuiteDir(suiteName)

	s := &model.Suite{}

	bytes, err := ioutil.ReadFile(suitePath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, s)

	if err != nil {
		return err
	}

	suite.RunSuite(*s)

	return nil
}

func isJsonFile(file os.FileInfo) bool {
	return !file.IsDir() && strings.HasSuffix(file.Name(), ".json")
}
