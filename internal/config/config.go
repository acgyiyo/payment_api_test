package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	Dbname   string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"DB_PORT"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("local") //set environment to use
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadSecretsFromSecretsService() {
	//TODO implement a service that load secrets from a secrets like user and password from a
	// service like AWS Secrets Manager or GCP Secret Manager
}
