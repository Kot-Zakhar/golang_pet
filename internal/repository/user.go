package repository

import (
	"context"
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
	query := `SELECT id, name, login, password, createdAt FROM users`
	rows, err := repo.db.Query(context, query)

	if err != nil {
		return nil, fmt.Errorf("UserRepository:GetAll:db.Query - %w", err)
		// TODO use this strategy across all errors
	}

	var users []model.User

	for rows.Next() && err == nil {
		var user model.User
		err = rows.Scan(&user.Id, &user.Name, &user.Login, &user.Password, &user.CreatedAt)
		users = append(users, user)
	}

	if err != nil {
		return nil, err
	} else {
		return &users, nil
	}
}
