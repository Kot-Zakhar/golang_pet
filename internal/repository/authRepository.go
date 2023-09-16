package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kot-zakhar/golang_pet/internal/model"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) InsertSession(context context.Context, session model.UserSession) (newSession model.UserSession, err error) {
	query := `
		INSERT INTO userSessions
		(userId, refreshToken, userAgent, fingerprint, expiresAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, userId, refreshToken, userAgent, fingerprint, expiresAt, createdAt`

	err = repo.db.QueryRow(context, query,
		session.UserId,
		session.RefreshToken,
		session.UserAgent,
		session.Fingerprint,
		session.ExpiresAt,
		session.CreatedAt,
	).Scan(
		&newSession.Id,
		&newSession.UserId,
		&newSession.RefreshToken,
		&newSession.UserAgent,
		&newSession.Fingerprint,
		&newSession.ExpiresAt,
		&newSession.CreatedAt,
	)

	if err != nil {
		return session, fmt.Errorf("AuthRepository:CreateSession - %w", err)
	}

	return newSession, nil
}
