package run

import (
	"encoding/json"
	"github.com/vusalalishov/api-tester/core/log"
	"github.com/vusalalishov/api-tester/core/model"
	"github.com/vusalalishov/api-tester/core/suite"
	"io/ioutil"
	"os"
	"strings"
)

func AllSuites(baseDir string) ([]*log.TestSuite, error) {
	fileInfos, err := ioutil.ReadDir(baseDir)

	if err != nil {
		return nil, err
	}

	var logs = make([]*log.TestSuite, 0)

	for _, fileInfo := range fileInfos {
		if isJsonFile(fileInfo) {
			suiteLog, err := Suite(baseDir, fileInfo.Name())
			if err != nil {
				return logs, err
			}
			logs = append(logs, suiteLog)
		}
	}

	return logs, nil
}

func Suite(baseDir string, suiteName string) (*log.TestSuite, error) {

	suitePath := baseDir + "/" + suiteName

	s := &model.Suite{}

	bytes, err := ioutil.ReadFile(suitePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, s)

	if err != nil {
		return nil, err
	}

	return suite.RunSuite(*s), nil
}

func isJsonFile(file os.FileInfo) bool {
	return !file.IsDir() && strings.HasSuffix(file.Name(), ".json")
}
