package models

type Config struct {
	Url          string `yaml:"url"`
	Auth         *Auth  `yaml:"auth"`
	Token        string `yaml:"token"`
	RefreshToken string `yaml:"refresh_token"`
}

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}