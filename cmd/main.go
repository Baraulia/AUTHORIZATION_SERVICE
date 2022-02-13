package main

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/handler"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/pkg/database"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/repository"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/service"
)

func main() {
	logger := logging.GetLogger()
	db, err := database.NewPostgresDB(database.PostgresDB{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "secret",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logger.Panicf("failed to initialize db:%s", err.Error())
	}

	rep := repository.NewRepository(db, logger)
	ser := service.NewService(rep, logger)
	handlers := handler.NewHandler(ser, logger)
	// Setup router
	router := handlers.InitRoutes()
	port := "8080"
	logger.Fatal(router.Run(":" + port))
}
