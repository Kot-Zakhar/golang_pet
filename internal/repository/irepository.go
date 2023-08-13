package repository

import (
	"context"
)

type IRepository[E any, K any] interface {
	GetAll(context.Context) (*[]E, error)
	GetById(context.Context, K) (*[]E, error)
	Save(context.Context, *E) error
	Delete(context.Context, K) error
}
