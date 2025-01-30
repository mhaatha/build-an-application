package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPlayerStore struct {
	db *pgxpool.Pool
}

func NewPostgresPlayerStore(db *pgxpool.Pool) *PostgresPlayerStore {
	return &PostgresPlayerStore{db: db}
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	var score int

	query := `SELECT score from players WHERE name = $1`
	err := p.db.QueryRow(context.Background(), query, name).Scan(&score)
	if err != nil {
		log.Printf("error fetching player score: %v\n", err)
		return 0
	}
	return score
}

func (p *PostgresPlayerStore) RecordWin(name string) {
	query := `
		INSERT INTO players (name, score)
		VALUES ($1, 1)
		ON CONFLICT (name)
		DO UPDATE SET score = players.score + 1
	`

	_, err := p.db.Exec(context.Background(), query, name)
	if err != nil {
		log.Printf("error recording win: %v\n", err)
	}
}
