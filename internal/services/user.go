package service

import (
	"crypto/rand"
	"encoding/hex"

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

	return generateSecureToken(128), nil

}

func generateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
