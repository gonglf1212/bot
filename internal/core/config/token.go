package config

type Token struct {
	AppID       uint64 `yaml:"appid"`
	AccessToken string `yaml:"token"`
	Type        string `yaml:"type"`
}
