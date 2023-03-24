package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	service "github.com/qudecim/password-manager-backend/internal/services"
)

// Authentification
// Check param email
// Return enycription data
func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	d := service.Auth()
	out(w, d)
}

// Authentification
// Get email and denycrypt data
// Return token
func Authentification(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	d := service.Authentification()
	out(w, d)
}

// Registration
// We check params email and public key
// Return token
func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	d := service.Register()
	out(w, d)
}
