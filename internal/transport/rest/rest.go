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
	router.POST("/register/", rest.Register)

	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
