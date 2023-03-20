package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"http_port"`
	URL  string `mapstructure:"db_url"`
}

// HTTP struct {
// }
// DB struct {
// }

func NewConfig() (*Config, error) {
	viper.SetDefault("http_port", "80")
	viper.SetDefault("db_url", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	viper.AutomaticEnv()
	cfg := &Config{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
