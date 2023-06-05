package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func CreateCockroachDBConnection(dsn string) (*pgx.Conn, error) {
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, errors.New("Error parsing connection string:" + err.Error())
	}
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.New("Error passing connection config:" + err.Error())
	}

	return conn, nil
}
