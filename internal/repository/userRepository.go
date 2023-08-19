package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/kot-zakhar/golang_pet/internal/model"
)

type IUserRepository interface {
	IRepository[model.User, uint64]
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) GetAll(context context.Context) (*[]model.User, error) {
	query := `SELECT id, name, login, email, password, createdAt FROM users`
	rows, err := repo.db.Query(context, query)
	if err != nil {
		return nil, fmt.Errorf("UserRepository:GetAll:db.Query - %w", err)
	}

	var users []model.User = make([]model.User, 0)

	for rows.Next() && err == nil {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Login, &user.Email, &user.Password, &user.CreatedAt)
		users = append(users, user)
	}

	if err != nil {
		return nil, fmt.Errorf("UserRepository:GetAll:rows.Scan - %w", err)
	} else {
		return &users, nil
	}
}

func (repo *UserRepository) GetById(context context.Context, id uint64) (*model.User, error) {
	query := `
		SELECT id, name, login, email, password, createdAt
		FROM users
		WHERE id = $1`
	row := repo.db.QueryRow(context, query, id)

	var user model.User

	err := row.Scan(&user.Id, &user.Name, &user.Login, &user.Email, &user.Password, &user.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("UserRepository:GetById:row.Scan - %w", err)
	} else {
		return &user, nil
	}
}

func (repo *UserRepository) GetByLogin(context context.Context, login string) (*model.User, error) {
	query := `
		SELECT login
		FROM users
		WHERE login = $1`
	row := repo.db.QueryRow(context, query, login)

	var user model.User

	err := row.Scan(&user.Id, &user.Name, &user.Login, &user.Password, &user.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("UserRepository:GetByLogin:row.Scan - %w", err)
	} else {
		return &user, nil
	}
}

func (repo *UserRepository) Insert(context context.Context, user *model.User) (uint64, error) {
	query := `
		INSERT INTO users
		(name, login, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	err := repo.db.QueryRow(context, query, user.Name, user.Login, user.Email, user.Password).Scan(&user.Id)

	if err != nil {
		return 0, fmt.Errorf("UserRepository:Insert:row.Scan - %w", err)
	} else {
		return user.Id, nil
	}
}

func (repo *UserRepository) Update(context context.Context, id uint64, user *model.User) error {
	query := `
		UPDATE users
		SET
			name = $1,
			login = $2,
			email = $3,
			password = $4
		WHERE
			id = $5`
	_, err := repo.db.Exec(context, query, user.Name, user.Login, user.Email, user.Password, id)

	if err != nil {
		return fmt.Errorf("UserRepository:Insert:row.Scan - %w", err)
	}

	return nil
}

func (repo *UserRepository) Delete(context context.Context, id uint64) error {
	query := `
		DELETE FROM users
		where id = $1
	`

	_, err := repo.db.Exec(context, query, id)

	if err != nil {
		return fmt.Errorf("UserRepository:Delete:db.Exec - %w", err)
	}

	return nil

}
