package main

import (
	"encoding/json"
	"github.com/vusalalishov/api-tester/core/model"
	"github.com/vusalalishov/api-tester/core/suite"
	"io/ioutil"
)

func main() {

	s := &model.Suite{}

	bytes, err := ioutil.ReadFile("resources/suite-1.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, s)

	if err != nil {
		panic(err)
	}

	suite.RunSuite(*s)

}
