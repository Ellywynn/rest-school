package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() error {
	// Configure .env file
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Can't load .env file: %s\n", err.Error())
	}

	// Configure viper config
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
