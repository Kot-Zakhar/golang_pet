package model

import "time"

type User struct {
	Id           uint64
	Name         string
	Login        string
	PasswordHash []byte
	Salt         []byte
	Email        string
	CreatedAt    time.Time
}
