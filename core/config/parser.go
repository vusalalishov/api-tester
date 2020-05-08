package config

import (
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func InitConfig() error {
	cfg, err := parseArgs()
	if err != nil {
		return err
	} else {
		initialize(cfg)
		return nil
	}
}

func parseArgs() (*RunConfiguration, error) {
	parser := argparse.NewParser("rapit", "REST API tester")

	baseDir := parser.String("d", "directory", &argparse.Options{
		Required: true,
		Validate: func(args []string) error {
			if len(args) != 1 {
				return errors.New("only 1 directory should be specified")
			}
			finfo, err := os.Stat(args[0])
			if err != nil {
				return err
			}
			if !finfo.IsDir() {
				return errors.New(fmt.Sprintf("%s is not a directory", args[0]))
			}
			return nil
		},
		Help: "Directory containing suites, templates and scripts",
	})

	suite := parser.String("s", "suite", &argparse.Options{
		Required: false,
		Help:     "Label of the suite to run (multiple suites might have the same label)",
	})

	err := parser.Parse(os.Args)

	if err != nil {
		return nil, errors.New(parser.Usage(err))
	} else {
		return &RunConfiguration{
			BaseDir: *baseDir,
			Suite:   suite,
		}, nil
	}
}
