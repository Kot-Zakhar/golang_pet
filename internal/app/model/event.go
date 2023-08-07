package model

import "time"

type Event struct {
	Id          uint
	Title       string
	Description string
	CreatedAt   time.Time
}
