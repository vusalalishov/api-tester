package main

import (
	"github.com/vusalalishov/rapit/core/config"
	"github.com/vusalalishov/rapit/core/run"
)

func main() {

	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	if err := run.AllSuites(); err != nil {
		panic(err)
	}
}
