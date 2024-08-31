package main

import (
	"context"
	todolist "github.com/TauAdam/todo-list"
	"github.com/TauAdam/todo-list/pkg/handler"
	"github.com/TauAdam/todo-list/pkg/repository"
	"github.com/TauAdam/todo-list/pkg/service"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Port string   `yaml:"port"`
	DB   ConfigDB `yaml:"db"`
}

type ConfigDB struct {
	Host     string `yaml:"host"`
	UserName string `yaml:"username"`
	DbName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	Port     string `yaml:"port"`
	Password string `env:"DATABASE_PASSWORD"`
}

//	@title			Todo list API
//	@version		1.0
//	@description	This is a simple todo list API
//	@host			localhost:8080
//	@BasePath		/
//	@schemes		http

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	//if err := godotenv.Load(); err != nil {
	//	logrus.Fatalf("failed to load env variables: %v", err.Error())
	//}

	var cfg Config
	err := cleanenv.ReadConfig("config.yaml", &cfg)
	if err != nil {
		logrus.Fatalf("failed to read config: %v", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		DBName:   cfg.DB.DbName,
		SSLMode:  cfg.DB.SSLMode,
		Username: cfg.DB.UserName,
		Password: cfg.DB.Password,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todolist.Server)
	port := cfg.Port
	go func() {
		if err := server.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Failed to start server: %v", err)
		}
	}()
	logrus.Println("Todo list app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Todo list app shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %v", err)
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close: %v", err)
	}
}
