package main

import (
	"log"

	"github.com/ivanov-s-tmn/rest-api"
	handler "github.com/ivanov-s-tmn/rest-api/pkg/handlers"
)

func main() {
	handlers := new(handler.Handler)
	server := new(rest.Server)

	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
