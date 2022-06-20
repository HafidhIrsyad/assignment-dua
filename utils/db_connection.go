package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresPool(host, port, user, password, name string) (*pgxpool.Pool, error) {
	connCfg := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)
	return pgxpool.Connect(context.Background(), connCfg)
}
