package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qudecim/password-manager-backend/pkg/repository"
	"github.com/qudecim/password-manager-backend/pkg/service"
	"github.com/qudecim/password-manager-backend/pkg/transport"
	"github.com/qudecim/password-manager-backend/pkg/transport/rest"
	"github.com/qudecim/password-manager-backend/pkg/transport/socket"

	_ "github.com/go-sql-driver/mysql"
)

// Let's go mthfck
func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)

	router := gin.New()

	restHandler := rest.NewHandler(services)
	router = restHandler.InitRoutes(router)

	socketHandler := socket.NewHandler(services)
	router = socketHandler.InitRoutes(router)

	srv := new(transport.Server)
	if err := srv.Run("8080", router); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
