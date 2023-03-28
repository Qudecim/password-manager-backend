package service

import (
	"fmt"
	"math/big"

	"github.com/qudecim/password-manager-backend/internal/enycrypt"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
)

type ConfirmationStringResponse struct {
	ConfirmationString []byte
}

// It enycrypt some data and return
func Auth(email string) (*ConfirmationStringResponse, error) {

	var resp ConfirmationStringResponse

	userModel := mysql.UserModel{DB: DB}

	hasUser, err := userModel.Has(email)
	if !hasUser {
		var err error
		return nil, err
	}

	ConfiramtionString := enycrypt.RandStringRunes(12)
	_, err = userModel.SetConfiramtionString(email, ConfiramtionString)
	if err != nil {
		fmt.Println("Don't write cofirmation string")
	}

	user, _ := userModel.Get(email)

	n := new(big.Int)
	nd, _ := n.SetString(user.PublicKey, 10)

	var publicKey enycrypt.PublicKey
	publicKey.E = big.NewInt(65537)
	publicKey.N = nd

	en, err := enycrypt.EncryptRSA(&publicKey, []byte(ConfiramtionString))
	if err != nil {
		fmt.Println(err)
	}

	resp.ConfirmationString = en

	return &resp, err
}

// Get email and deenycrypt data
// We check it and create token
func Authentification() [6]int {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	return primes
}

func Gen(email string) (*enycrypt.PublicKey, error) {

	userModel := mysql.UserModel{DB: DB}

	hasUser, err := userModel.Has(email)
	if !hasUser {
		var err error
		return nil, err
	}

	ConfiramtionString := enycrypt.RandStringRunes(12)
	userModel.SetConfiramtionString(email, ConfiramtionString)

	pub, _, _ := enycrypt.GenerateKeys(1024)

	return pub, err
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
