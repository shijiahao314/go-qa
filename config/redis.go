package config

type RedisConfig struct {
	DB            int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password      string `mapstructure:"password" json:"password" yaml:"password"`
	ConnectionNum int    `mapstructure:"connection-num" json:"connection-num" yaml:"connection-num"`
}
