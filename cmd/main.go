package main

import (
	"log"

	server "github.com/mrnkslv/user-segmentation-service"
	"github.com/mrnkslv/user-segmentation-service/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
