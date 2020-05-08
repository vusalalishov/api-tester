package config

type RunConfiguration struct {
	BaseDir string
	Suite   *string
}

var runConfig *RunConfiguration

func initialize(cfg *RunConfiguration) {
	runConfig = cfg
}

func SuiteDir(suiteFile string) string {
	return GetSuiteDir() + "/" + suiteFile
}

func GetSuiteDir() string {
	return runConfig.BaseDir + "/suites"
}

func ScriptDir(scriptFile string) string {
	return runConfig.BaseDir + "/scripts/" + scriptFile
}

func PayloadDir(file string) string {
	return runConfig.BaseDir + "/payloads/" + file
}
