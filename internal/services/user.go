package service

import (
	"fmt"

	"github.com/qudecim/password-manager-backend/internal/enycrypt"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
)

// It enycrypt some data and return
func Auth(email string) (string, error) {

	userModel := mysql.UserModel{DB: DB}

	hasUser, err := userModel.Has(email)
	if hasUser {
		var err error
		return "", err
	}

	ConfiramtionString := enycrypt.RandStringRunes(12)
	userModel.SetConfiramtionString(email, ConfiramtionString)

	return "nil", err
}

// Get email and deenycrypt data
// We check it and create token
func Authentification() [6]int {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	return primes
}

// It write email and public method in DB
// after that it is creating token
func Register(email string, publicKey string) (bool, error) {

	userModel := mysql.UserModel{DB: DB}

	fmt.Println(email)

	hasUser, err := userModel.Has(email)
	if err != nil {
		return false, err
	}

	if hasUser {
		var err error
		return false, err
	}

	success, err := userModel.Add(email, publicKey)

	return success, nil
}

// Get token and check it
func checkToken() {

}

// Create token
func createToken() {

}
