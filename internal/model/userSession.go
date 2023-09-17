package model

import (
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	Id           int
	UserId       int
	RefreshToken uuid.UUID
	UserAgent    string
	// TODO-Q: how do we actually check the validity of fingerprint
	Fingerprint string
	ExpiresAt   time.Time
	CreatedAt   time.Time
}
