package main

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/handler"
	"github.com/TauAdam/todo-list/pkg/repository"
	"github.com/TauAdam/todo-list/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todolist.Server)
	port := "8080"
	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
