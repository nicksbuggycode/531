package db

import (
	"context"
	"database/sql"
	"fmt"
)

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

// execute a function within a db transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
tx, err := store.db.BeginTx(ctx, nil)
if err != nil{
	return err
}
q := New(tx)
err = fn(q)
if err != nil{
	if rbErr := tx.Rollback(); rbErr != nil{
		return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
	}
	return err
}
return  tx.Commit()

}