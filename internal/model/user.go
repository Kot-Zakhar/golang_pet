package model

import "time"

type User struct {
	Id           int
	Name         string
	Login        string
	PasswordHash []byte
	Salt         []byte
	Email        string
	CreatedAt    time.Time
}
