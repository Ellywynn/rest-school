package server

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ellywynn/rest-school/handler"
	"github.com/ellywynn/rest-school/repository"
	"github.com/ellywynn/rest-school/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	httpServer *http.Server
}

func NewApp() *Server {
	db, err := repository.NewPostgres(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBUser:   viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Database connection error: %s\n", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	port := viper.GetString("server.port")

	server, err := createServer(port, handler.InitRoutes())
	if err != nil {
		logrus.Fatalf("Cannot create the server: %s", err.Error())
	}

	return server
}

func (s *Server) Run(port string) error {
	return s.httpServer.ListenAndServe()
}

func createServer(port string, handler http.Handler) (*Server, error) {
	// Get server configuration from config file
	headerBytes, err := strconv.Atoi(viper.GetString("server.maxHeaderBytesMB"))
	if err != nil {
		return nil, err
	}

	writeTimeout, err := strconv.Atoi(viper.GetString("server.writeTimeoutSeconds"))
	if err != nil {
		return nil, err
	}

	readTimeout, err := strconv.Atoi(viper.GetString("server.readTimeoutSeconds"))
	if err != nil {
		return nil, err
	}

	// Configure and create http server
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: headerBytes << 20,
			WriteTimeout:   time.Duration(writeTimeout) * time.Second,
			ReadTimeout:    time.Duration(readTimeout) * time.Second,
		},
	}, nil
}
