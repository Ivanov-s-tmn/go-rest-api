package main

import (
	"log"

	"github.com/ivanov-s-tmn/go-rest-api"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/handlers"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/repository"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	srvc := service.NewService(repo)
	handlers := handlers.NewHandler(srvc)

	server := new(rest.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
