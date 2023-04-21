package models

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"requred"`
	Email    string `json:"username" binding:"requred"`
	Password string `json:"password" binding:"requred"`
}
