package rest

import (
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	rest "github.com/qudecim/password-manager-backend/internal/transport/rest/controllers"
)

var addr = flag.String("addr", ":8080", "http service address")

func Run() {
	router := httprouter.New()

	router.POST("/auth/", rest.Auth)
	router.POST("/auth_confirm/", rest.AuthConfirm)
	router.POST("/register/", rest.Register)

	router.GET("/device", rest.Auth)
	router.POST("/device", rest.Auth)
	router.DELETE("/device/:id", rest.Auth)

	router.POST("/send/:id", rest.Auth)

	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
