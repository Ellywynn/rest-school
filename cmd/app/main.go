package main

import (
	"log"

	"github.com/ellywynn/rest-school/pkg/config"
	"github.com/ellywynn/rest-school/pkg/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := config.Init(); err != nil {
		logrus.Fatalf("Config error: %s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("server.port")); err != nil {
		log.Fatalf("Error while running the server: %s", err.Error())
	}
}
