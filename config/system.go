package config

type SystemConfig struct {
	AppName       string  `yaml:"app_name"`
	Version       float32 `yaml:"version"`
	GitHookUrl    string  `yaml:"githook_url"`
	AppRepository string  `yaml:"app_repository"`
}
