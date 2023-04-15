package models

type Secret struct {
	Id          int    `json:"id"`
	Enycryption string `json:"enycryption"`
}

type UsersSecret struct {
	Id       int
	UserId   int
	SecretId int
}
