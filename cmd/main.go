package main

import (
	"context"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/handler"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/pkg/database"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/repository"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/server"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := logging.GetLogger()
	db, err := database.NewPostgresDB(database.PostgresDB{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		logger.Panicf("failed to initialize db:%s", err.Error())
	}

	rep := repository.NewRepository(db, logger)
	ser := service.NewService(rep, logger)
	handlers := handler.NewHandler(ser, logger)

	port := os.Getenv("API_SERVER_PORT")
	serv := new(server.Server)

	go func() {
		err := serv.Run(port, handlers.InitRoutes())
		if err != nil {
			logger.Panicf("Error occured while running http server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := serv.Shutdown(context.Background()); err != nil {
		logger.Panicf("Error occured while shutting down http server: %s", err.Error())
	}
}
