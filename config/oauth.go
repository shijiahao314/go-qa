package config

type OAuthConfig struct {
	Github GithubConfig `mapstructure:"github" json:"github" yaml:"github"`
}

type GithubConfig struct {
	ClientID     string `mapstructure:"client-id" json:"client-id" yaml:"client-id"`
	ClientSecret string `mapstructure:"client-secret" json:"client-secret" yaml:"client-secret"`
	RedirectURL  string `mapstructure:"redirect-url" json:"redirect-url" yaml:"redirect-url"`
}
