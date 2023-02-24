package models

type Config struct {
	Url    string  `yaml:"url"`
	Auth   *Auth   `yaml:"auth"`
	Token  string  `yaml:"token"`
	Aliyun *Aliyun `yaml:"aliyun"`
	PikPak *PikPak `yaml:"pikpak"`
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
