package mysql

import (
	"context"
	"time"

	"github.com/qudecim/password-manager-backend/internal/app"
	"github.com/qudecim/password-manager-backend/internal/models"
)

func Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func Get(Email string) (*models.User, error) {
	var user models.User
	err := app.DB.QueryRow("select * from users where email = ?", Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserHas(user *models.User) (bool, error) {
	rows, err := app.DB.Query("select * from Users where email = ?", user.Email)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func UserAdd(user *models.User) error {
	_, err := app.DB.ExecContext(context.Background(), "INSERT INTO Users (email, password) VALUES (?,?)", user.Email, user.Password)
	return err
}

func UserAuth(user *models.User) (int, error) {
	var id int

	err := app.DB.QueryRow("select ID from Users where email = ? and password = ?", user.Email, user.Password).Scan(id)
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

func TokenAdd(user *models.User, token string) error {
	date := time.Now().AddDate(1, 0, 0).Format(time.RFC3339)
	_, err := app.DB.ExecContext(context.Background(), "INSERT INTO Tokens (user_id, token, expired_at) VALUES (?,?,?)", user.ID, token, date)
	return err
}
