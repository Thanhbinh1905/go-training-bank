package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL         string        `mapstructure:"POSTGRES_URL"`
	Production          bool          `mapstructure:"PRODUCTION"`
	TokenSymmetricKet   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.SetConfigName(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
