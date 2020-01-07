package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	UserConfig   `yaml:",inline"`
	SystemConfig `yaml:",inline"`
}

var Cfg Config
var CurrentDir string

func init() {
	var pwdErr error
	CurrentDir, pwdErr = os.Getwd()
	if pwdErr != nil {
		panic(pwdErr)
	}

	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	//jsonErr := json.Unmarshal(configFile, &Cfg)
	//if jsonErr != nil {
	//	panic(err)
	//}

	if err := yaml.Unmarshal(configFile, &Cfg); err != nil {
		panic(err)
	}

	if "" == Cfg.DashboardEntrance || !strings.HasPrefix(Cfg.DashboardEntrance, "/") {
		Cfg.DashboardEntrance = "/admin"
	}

	Cfg.AppName = "Go"
	Cfg.Version = 2.2
	Cfg.GitHookUrl = "/api/git_push_hook"
	Cfg.AppRepository = ""
}
