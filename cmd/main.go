package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ivanov-s-tmn/go-rest-api"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/handlers"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/repository"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error reading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error connecting db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	handlers := handlers.NewHandler(srvc)

	server := new(rest.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error while running server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("Error while shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("Error while closing db connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
