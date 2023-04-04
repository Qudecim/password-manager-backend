package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/qudecim/password-manager-backend/internal/models"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
	service "github.com/qudecim/password-manager-backend/internal/services"
)

type authResponse struct {
	token string
}

// Authentification
func Auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user models.User

	// Получаем user из json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	// Проверяем существование
	hasUser, err := mysql.UserHas(&user)
	if err != nil {
		out(w, false, nil, err)
		return
	}
	if !hasUser {
		err := errors.New("User not exist")
		out(w, false, nil, err)
		return
	}

	// Аутентификация
	err = service.Auth(&user)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	// Создание токена
	token, err := service.CreateToken(&user)

	answer := authResponse{token}

	out(w, true, answer, err)
}

// Registration
func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var user models.User

	// Получаем user из json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	// Проверяем существование
	hasUser, err := mysql.UserHas(&user)
	if err != nil {
		out(w, false, nil, err)
		return
	}
	if hasUser {
		out(w, false, nil, err)
		return
	}

	// Добавлени пользователя
	err = service.Register(&user)
	if err != nil {
		out(w, false, nil, err)
		return
	}

	out(w, true, nil, err)
}
