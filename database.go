package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	dbPassowrd, err := GetDBPassword("./.env")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "notrhel", dbPassowrd, "localhost", "5432", "scoring")

	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return dbPool, nil
}
