package models

type Secret struct {
	Id          int    `json:"id" db:"id"`
	UserId      int    `json:"user_id" db:"user_id"`
	Enycryption string `json:"enycryption" binding:"required" db:"enycryption"`
}
