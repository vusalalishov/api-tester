package main

import (
	model "api-tester/core/model"
	suite "api-tester/core/suite"
	"encoding/json"
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
