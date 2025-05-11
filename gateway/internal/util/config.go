package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment     string `mapstructure:"ENVIRONMENT"`
	ServerAddress   string `mapstructure:"SERVER_ADDRESS"`
	RateLimitEnable bool   `mapstructure:"RATE_LIMIT_ENABLE"`
	RateLimitRps    int    `mapstructure:"RATE_LIMIT_RPS"`
	RateLimitBurst  int    `mapstructure:"RATE_LIMIT_BURST"`
	TokenSecret     string `mapstructure:"TOKEN_SECRET_KEY"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	var config Config

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}
