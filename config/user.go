package config

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
	DashboardEntrance string   `yaml:"dashboardEntrance"`
}
