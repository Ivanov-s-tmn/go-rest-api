package main

import (
	"log"

	"github.com/ivanov-s-tmn/go-rest-api"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/handlers"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/repository"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5433",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("error connecting db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	handlers := handlers.NewHandler(srvc)

	server := new(rest.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
