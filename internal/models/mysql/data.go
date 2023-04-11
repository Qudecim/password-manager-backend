package mysql

import (
	"context"

	"github.com/qudecim/password-manager-backend/internal/app"
	"github.com/qudecim/password-manager-backend/internal/models"
)

func DataAdd(data models.Data) error {
	_, err := app.DB.ExecContext(context.Background(), "INSERT INTO Data (user_id, data) VALUES (?,?)", data.UserId, data.Data)
	return err
}

func DataGet(user *models.User) ([]models.Data, error) {

	var dataSlice []models.Data

	rows, err := app.DB.Query("select id, data from Data where user_id = ?", user.ID)
	if err != nil {
		return dataSlice, err
	}

	for rows.Next() {
		var data models.Data
		if err := rows.Scan(&data.ID, &data.Data); err != nil {
			return dataSlice, err
		}
		dataSlice = append(dataSlice, data)
	}

	return dataSlice, nil
}
