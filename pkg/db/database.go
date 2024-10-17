package db

import (
	"context"
	"hype-casino-platform/pkg/kgserr"
)

// Database is an interface that defines the methods required to interact with a database.
type Database interface {
	// GetConn returns a connection to the database.
	GetConn(ctx context.Context) any
	// GetTx returns a transaction from current context. If not found, return nil
	GetTx(ctx context.Context) any
	// Begin starts a transaction with current context. If the transaction is already started, return an error
	Begin(ctx context.Context) (context.Context, *kgserr.KgsError)
	// Commit commits the transaction with current context. If the transaction is not started, return an error
	Commit(ctx context.Context) (context.Context, *kgserr.KgsError)
	// Rollback rolls back the transaction with current context. If the transaction is not started, return an error
	Rollback(ctx context.Context) (context.Context, *kgserr.KgsError)
}
