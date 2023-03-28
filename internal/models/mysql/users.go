package mysql

import (
	"context"
	"database/sql"

	"github.com/qudecim/password-manager-backend/internal/models"
)

// SnippetModel - Определяем тип который обертывает пул подключения sql.DB
type UserModel struct {
	DB *sql.DB
}

// Insert - Метод для создания новой заметки в базе дынных.
func (m *UserModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *UserModel) Get(Email string) (*models.User, error) {
	var user models.User
	err := m.DB.QueryRow("select * from Users where Email = ?", Email).Scan(&user.ID, &user.Email, &user.PublicKey, &user.ConfirmationString)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Is exsist user
func (m *UserModel) Has(Email string) (bool, error) {
	rows, err := m.DB.Query("select * from Users where Email = ?", Email)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

// Add user
func (m *UserModel) Add(Email string, PublicKey string) (bool, error) {
	_, err := m.DB.ExecContext(context.Background(), "INSERT INTO Users (Email, PublicKey) VALUES (?,?)", Email, PublicKey)
	return true, err
}

// Set Confiramtion String don't encoded
func (m *UserModel) SetConfiramtionString(Email string, ConfirmationString string) (bool, error) {
	_, err := m.DB.Exec("UPDATE Users SET ConfirmationString = ? WHERE Email = ?", ConfirmationString, Email)
	return true, err
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (m *UserModel) Latest() ([]*models.User, error) {
	return nil, nil
}
