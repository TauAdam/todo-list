package main

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	server := new(todolist.Server)
	port := "8080"
	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
