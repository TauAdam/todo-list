package main

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/handler"
	"github.com/TauAdam/todo-list/pkg/repository"
	"github.com/TauAdam/todo-list/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to init config: %v", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env variables: %v", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}
	repos := repository.NewRepository(db)
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
