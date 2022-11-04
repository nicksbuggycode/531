package db

import "database/sql"

// store provides all functions to execute db queries (individually) and transactions (batches of queries - ACID principle) -- we already have the query struct, embedding it in the store struct is composition, we keep and extend its functionality
type Store struct {
	*Queries
	db *sql.DB
}

// Create new Store object
func NewStore(db *sql.DB) *Store{
	return &Store{
		db: db,
		Queries: New(db),
	}
}