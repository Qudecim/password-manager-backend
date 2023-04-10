package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	match_rand "math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/qudecim/password-manager-backend/internal/models"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
)

type ConfirmationStringResponse struct {
	ConfirmationString []byte
}

func Auth(user *models.User) error {

	id, err := mysql.UserAuth(user)
	user.ID = id

	return err
}

func Register(user *models.User) error {

	err := mysql.UserAdd(user)

	return err
}

func checkToken() {

}

func CreateToken(user *models.User) (string, error) {

	// Рандомное число
	rand_int := match_rand.Int()

	// Строка токена
	token := strconv.Itoa(user.ID) + "|" + strconv.Itoa(rand_int)

	// sha128
	a := sha256.New()
	a.Write([]byte(token))
	hash := hex.EncodeToString(a.Sum(nil))

	err := mysql.TokenAdd(user, hash)

	return token, err

}

func IsAuth(r *http.Request) bool {

	token, err := getToken(r)
	if err != nil {
		return false
	}

	// Чекаем
	hasToken, _ := mysql.TokenHas(token)

	return hasToken
}

func GetUser(r *http.Request) (*models.User, error) {

	token, err := getToken(r)
	if err != nil {
		return nil, err
	}

	user, err := mysql.GetByToken(token)

	return user, err

}

func getToken(r *http.Request) (string, error) {

	reqToken := r.Header.Get("Authorization")

	if reqToken == "" {
		err := errors.New("Token didn't find")
		return "", err
	}

	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Хешируем
	a := sha256.New()
	a.Write([]byte(reqToken))
	hash := hex.EncodeToString(a.Sum(nil))

	return hash, nil

}
