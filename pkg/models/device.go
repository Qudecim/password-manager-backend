package models

type Device struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Uid       string `json:"uid" db:"uid"`
	PublicKey string `json:"public_key" db:"public_key"`
}
