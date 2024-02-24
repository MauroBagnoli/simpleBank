package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// tx, err := store.db.
// }
