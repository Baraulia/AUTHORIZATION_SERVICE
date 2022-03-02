package main

import (
	"context"
	"os"
	"os/signal"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC/grpcServer"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/handler"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/database"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/repository"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/server"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/service"
	"syscall"
)

// @title Authorization Service
// @description Authorization Service for Food Delivery Application
func main() {
	logger := logging.GetLogger()
	db, err := database.NewPostgresDB(database.PostgresDB{
		Host:     "159.223.1.135",
		Port:     "5433",
		Username: "authorizeteam1",
		Password: "qwerty",
		DBName:   "authorize_db",
		SSLMode:  "disable",
	})
	if err != nil {
		logger.Panicf("failed to initialize db:%s", err.Error())
	}

	rep := repository.NewRepository(db, logger)
	ser := service.NewService(rep, logger)
	handlers := handler.NewHandler(ser, logger)

	port := "8080"
	serv := new(server.Server)

	service.Secret = "ygKG2872@gk&GF26VDEWLsfret23#qw"

	go func() {
		err := serv.Run(port, handlers.InitRoutes())
		if err != nil {
			logger.Panicf("Error occured while running http server: %s", err.Error())
		}
	}()

	go func() {
		grpcServer.NewGRPCServer(ser)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := serv.Shutdown(context.Background()); err != nil {
		logger.Panicf("Error occured while shutting down http server: %s", err.Error())
	}
}
