package main

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/handler"
	"github.com/TauAdam/todo-list/pkg/repository"
	"github.com/TauAdam/todo-list/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to init config: %v", err)
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todolist.Server)
	port := viper.GetString("port")
	if err := server.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
