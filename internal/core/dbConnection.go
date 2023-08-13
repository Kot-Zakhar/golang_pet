package core

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectPgx connects to the pgxpool
func ConnectPgx(connectionString string) (*pgxpool.Pool, error) {
	cfgPgx, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, err
	}
	dbpool, err := pgxpool.NewWithConfig(context.Background(), cfgPgx)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}
