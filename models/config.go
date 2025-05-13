package models

// 主配置结构体
type Config struct {
	Url         string       `yaml:"url"`
	Auth        *Auth        `yaml:"auth"`
	Token       string       `yaml:"token"`
	Aliyun      *Aliyun      `yaml:"aliyun"`
	PikPak      *PikPak      `yaml:"pikpak"`
	OneDriveApp *OneDriveApp `yaml:"onedrive_app"`
}

// Alist 登录认证信息
type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// 阿里云盘配置
type Aliyun struct {
	Enable       bool   `yaml:"enable"`
	RefreshToken string `yaml:"refresh_token"`
}

// PikPak 配置
type PikPak struct {
	Enable                bool   `yaml:"enable"`
	UseTranscodingAddress bool   `yaml:"use_transcoding_address"`
	Username              string `yaml:"username"`
	Password              string `yaml:"password"`
}

// OneDrive APP 配置
type OneDriveApp struct {
	Enable  bool         `yaml:"enable"`
	Region  string       `yaml:"region"`
	Tenants []TenantList `yaml:"tenants"`
}

// OneDrive APP 租户信息
type TenantList struct {
	Id           int    `yaml:"id"`
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	TenantId     string `yaml:"tenant_id"`
}
