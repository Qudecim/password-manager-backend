package models

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Email    string `json:"username"`
	Password string `json:"password"`
}
