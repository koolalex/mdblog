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

var (
	Cfg        Config
	CurrentDir string
)

func init() {
	var err error
	if CurrentDir, err = os.Getwd(); err == nil {
		if b, err := ioutil.ReadFile("config.yml"); err == nil {
			if err = yaml.Unmarshal(b, &Cfg); err == nil {
				return
			}
		}
	}
	panic(err)
}
