package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource            string        `mapstructure:"DB_SOURCE"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	Production          bool
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	config := &Config{
		DBSource:            viper.GetString("DB_SOURCE"),
		TokenSymmetricKey:   viper.GetString("TOKEN_SYMMETRIC_KEY"),
		AccessTokenDuration: viper.GetDuration("ACCESS_TOKEN_DURATION"),
	}

	if config.DBSource == "" {
		return nil, fmt.Errorf("DB_SOURCE not set")
	}

	return config, nil
}
