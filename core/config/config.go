package config

type RunConfiguration struct {
	BaseDir string
}

var runConfig *RunConfiguration

func Init(baseDir string) {
	runConfig = &RunConfiguration{
		BaseDir: baseDir,
	}
}

func GetBaseDir() string {
	return runConfig.BaseDir
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

func SuiteDir(file string) string {
	return runConfig.BaseDir + "/suites/" + file
}
