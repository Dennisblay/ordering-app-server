package database

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"ordering-server/internal/database/models"
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

//func (s *Store) Health() map[string]string {
//
//}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
//func (s *service) Close() error {
//
//}
