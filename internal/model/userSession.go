package model

import "time"

type UserSession struct {
	Id           int
	UserId       int
	RefreshToken string
	UserAgent    string
	Fingerprint  string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
