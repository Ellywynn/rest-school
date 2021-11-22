package main

import (
	"log"

	"github.com/ellywynn/rest-school/config"
	"github.com/ellywynn/rest-school/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := config.Init(); err != nil {
		logrus.Fatalf("Config error: %s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}
}
