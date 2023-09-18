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
		(userId, userAgent, fingerprint, expiresAt, createdAt)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, userId, refreshToken, userAgent, fingerprint, expiresAt, createdAt`

	err = repo.db.QueryRow(context, query,
		session.UserId,
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

func (repo *AuthRepository) DeleteSession(context context.Context, refreshToken string) error {
	query := `
		DELETE FROM userSessions
		WHERE refreshToken = $1`

	_, err := repo.db.Exec(context, query, refreshToken)

	if err != nil {
		return fmt.Errorf("AuthRepository:DeleteSession - %w", err)
	}

	return nil
}

func (repo *AuthRepository) GetAndDeleteSession(context context.Context, refreshToken string) (session model.UserSession, err error) {
	query := `
		DELETE FROM userSessions
		WHERE refreshToken = $1
		RETURNING id, userId, refreshToken, userAgent, fingerprint, expiresAt, createdAt`

	err = repo.db.QueryRow(context, query, refreshToken).Scan(
		&session.Id,
		&session.UserId,
		&session.RefreshToken,
		&session.UserAgent,
		&session.Fingerprint,
		&session.ExpiresAt,
		&session.CreatedAt,
	)

	if err != nil {
		return session, fmt.Errorf("AuthRepository:GetAndDeleteSession - %w", err)
	}

	return session, nil
}
