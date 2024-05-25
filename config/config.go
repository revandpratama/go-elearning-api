package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	PORT        string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_URL      string
}

var ENV *config

func LoadConfig() {

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}

	log.Println("config is running successfully")
}
