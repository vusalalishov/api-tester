package main

import (
	"encoding/json"
	"fmt"
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

	log := suite.RunSuite(*s, "resources/", "resources/scripts/")

	fmt.Println(log.Title)
	fmt.Println(log.Status)

	for _, testCase := range log.Cases {
		fmt.Println(testCase.Title)
		fmt.Println(testCase.Status)

		for _, scenario := range testCase.Scenarios {
			fmt.Println(scenario.Title)
			fmt.Println(scenario.Status)
			for _, message := range scenario.Messages {
				fmt.Println(message)
			}
		}

	}

}
