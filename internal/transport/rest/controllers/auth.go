package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/qudecim/password-manager-backend/internal/models"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
	service "github.com/qudecim/password-manager-backend/internal/services"
)

// Authentification
func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user models.User

	// Получаем user из json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверяем существование
	hasUser, err := mysql.UserHas(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !hasUser {
		http.Error(w, "This email is not using", http.StatusBadRequest)
		return
	}

	// Аутентификация
	err = service.Auth(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := service.CreateToken(&user)

	out(w, d)
}

// Registration
func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user models.User

	// Получаем user из json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверяем существование
	hasUser, err := mysql.UserHas(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if hasUser {
		http.Error(w, "This email is using", http.StatusBadRequest)
		return
	}

	// Добавлени пользователя
	d, err := service.Register(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out(w, d)
}
