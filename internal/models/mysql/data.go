package mysql

import (
	"context"

	"github.com/qudecim/password-manager-backend/internal/app"
	"github.com/qudecim/password-manager-backend/internal/models"
)

func DataAdd(user *models.User, data string) error {
	_, err := app.DB.ExecContext(context.Background(), "INSERT INTO Data (user_id, data) VALUES (?,?)", user.ID, data)
	return err
}

func DataGet(user *models.User) error {
	err := app.DB.QueryRow("select * from Data where user_id = ?", Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
