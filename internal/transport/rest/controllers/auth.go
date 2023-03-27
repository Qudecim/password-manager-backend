package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	service "github.com/qudecim/password-manager-backend/internal/services"
)

type User struct {
	id        int
	Email     string
	PublicKey string
}

// Authentification
// Check param email
// Return enycription data
func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d, err := service.Auth(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d, err := service.Register(user.Email, user.PublicKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out(w, d)
}
