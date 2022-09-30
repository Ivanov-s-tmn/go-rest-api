package main

import (
	"log"
	"os"

	"github.com/ivanov-s-tmn/go-rest-api"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/handlers"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/repository"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error reading env variables: %s", err.Error())
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
