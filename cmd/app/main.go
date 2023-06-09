package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qudecim/password-manager-backend/pkg/repository"
	"github.com/qudecim/password-manager-backend/pkg/service"
	"github.com/qudecim/password-manager-backend/pkg/transport"
	"github.com/qudecim/password-manager-backend/pkg/transport/rest"
	"github.com/qudecim/password-manager-backend/pkg/transport/socket"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
)

// Let's go mthfck
func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewMysqlDB(repository.Config{
		Host:     viper.GetString("db_host"),
		Port:     viper.GetString("db_port"),
		Username: viper.GetString("db_username"),
		Password: viper.GetString("db_password"),
		DBname:   viper.GetString("db_name"),
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	hub := socket.NewHub()
	go hub.Run()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	router := gin.New()

	restHandler := rest.NewHandler(services)
	router = restHandler.InitRoutes(router)

	socketHandler := socket.NewHandler(hub, services)
	router = socketHandler.InitRoutes(router)

	srv := new(transport.Server)
	if err := srv.Run(viper.GetString("port"), router); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("../../configs") // home/pm/configs
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
