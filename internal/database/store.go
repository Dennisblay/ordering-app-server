package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	db "ordering-server/pkg/models"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*db.Queries
	connPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) *Store {
	return &Store{
		connPool: connPool,
		Queries:  db.New(connPool),
	}
}
