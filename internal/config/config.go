package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBURL string `mapstructure:"DB_URL"`
}

func LoadConfig() Config {
	v := viper.New()

	//check os env vars
	v.AutomaticEnv()

	//bind the field DB_URL
	v.BindEnv("DB_URL")

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config %v\n", cfg)
	}

	if cfg.DBURL == "" {
		log.Fatalf("DB_URL has not been set in the .zshrc or .bashrc file")
	}

	return cfg
}
