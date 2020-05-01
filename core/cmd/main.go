package main

import (
	"github.com/vusalalishov/api-tester/core/config"
	"github.com/vusalalishov/api-tester/core/run"
)

func main() {
	config.Init("/Users/vusalalishov/projects/go/src/api-tester/tests")
	if err := run.AllSuites(); err != nil {
		panic(err)
	}
}
