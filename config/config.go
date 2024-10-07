package config

type Config struct {
	Port        int            `mapstructure:"port" json:"port" yaml:"port"`
	Token       string         `mapstructure:"token" json:"token" yaml:"token"`
	Salt        string         `mapstructure:"salt" json:"salt" yaml:"salt"`
	Database    DatabaseConfig `mapstructure:"database" json:"database" yaml:"database"`
	Redis       RedisConfig    `mapstructure:"redis" json:"redis" yaml:"redis"`
	OAuthConfig OAuthConfig    `mapstructure:"oauth" json:"oauth" yaml:"oauth"`
	// Etcd        EtcdConfig     `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
}
