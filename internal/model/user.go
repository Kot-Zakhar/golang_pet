package model

import "time"

type User struct {
	Id        uint64
	Name      string
	Login     string
	Password  string
	Email     string
	CreatedAt time.Time
}
