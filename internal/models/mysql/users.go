package mysql

import (
	"context"

	"github.com/qudecim/password-manager-backend/internal/app"
	"github.com/qudecim/password-manager-backend/internal/models"
)

func Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func Get(Email string) (*models.User, error) {
	var user models.User
	err := app.DB.QueryRow("select * from Users where Email = ?", Email).Scan(&user.ID, &user.Email, &user.PublicKey, &user.ConfirmationString)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserHas(user *models.User) (bool, error) {
	rows, err := app.DB.Query("select * from Users where Email = ?", user.Email)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func UserAdd(user *models.User) (bool, error) {
	_, err := app.DB.ExecContext(context.Background(), "INSERT INTO Users (Email, Password) VALUES (?,?)", user.Email, user.Password)
	return true, err
}

func UserAuth(user *models.User) (int, error) {
	var id int

	err := app.DB.QueryRow("select ID from Users where Email = ? and Password = ?", user.Email, user.Password).Scan(id)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		var err error
		err.Error()
		return 0, err
	}

	return id, nil
}
