package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	UserConfig   `yaml:"user"`
	SystemConfig `yaml:"system"`
}

type SystemConfig struct {
	AppName       string `yaml:"app_name"`
	Version       string `yaml:"version"`
	GitHookUrl    string `yaml:"githook_url"`
	AppRepository string `yaml:"app_repository"`
}

type UserConfig struct {
	SiteName          string   `yaml:"site_name"`
	SiteKeywords      string   `yaml:"site_keywords"`
	SiteDescription   string   `yaml:"site_description"`
	Author            string   `yaml:"author"`
	Icp               string   `yaml:"icp"`
	TimeLayout        string   `yaml:"time_layout"`
	Port              int      `yaml:"port"`
	WebHookSecret     string   `yaml:"webhook_secret"`
	UtterancesRepo    string   `yaml:"utterances_repo"`
	PageSize          int      `yaml:"page_size"`
	DescriptionLen    int      `yaml:"description_len"`
	DocumentPath      string   `yaml:"document_path"`
	CategoryDocNumber int      `yaml:"category_doc_number"`
	ThemeColor        string   `yaml:"theme_color"`
	ThemeOption       []string `yaml:"theme_option"`
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

	if err := yaml.Unmarshal(configFile, &Cfg); err != nil {
		panic(err)
	}

}
