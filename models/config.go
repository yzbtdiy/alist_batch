package models

type Config struct {
	Url         string       `yaml:"url"`
	Auth        *Auth        `yaml:"auth"`
	Token       string       `yaml:"token"`
	Aliyun      *Aliyun      `yaml:"aliyun"`
	PikPak      *PikPak      `yaml:"pikpak"`
	OneDriveApp *OneDriveApp `yaml:"onedrive_app"`
}

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Aliyun struct {
	Enable       bool   `yaml:"enable"`
	RefreshToken string `yaml:"refresh_token"`
}

type PikPak struct {
	Enable   bool   `yaml:"enable"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type OneDriveApp struct {
	Enable bool         `yaml:"enable"`
	Region string       `yaml:"region"`
	Tenant []TenantList `yaml:"tenants"`
}

type TenantList struct {
	Id           int    `yaml:"id"`
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	TenantId     string `yaml:"tenant_id"`
}
