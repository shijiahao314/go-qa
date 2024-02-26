package config

type OAuthGithub struct {
	ClientID     string `mapstructure:"client-id" json:"client-id" yaml:"client-id"`
	ClientSecret string `mapstructure:"client-secret" json:"client-secret" yaml:"client-secret"`
}

type Config struct {
	Port     int      `mapstructure:"port" json:"port" yaml:"port"`
	Token    string   `mapstructure:"token" json:"token" yaml:"token"`
	Salt     string   `mapstructure:"salt" json:"salt" yaml:"salt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	OAuthGithub
}
