package main

import (
	"github.com/c9llm3bones/Films-VK"
	"github.com/c9llm3bones/Films-VK/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(Films.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while while running http server : %s", err.Error())
	}

}
