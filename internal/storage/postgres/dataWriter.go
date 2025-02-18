package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"

	db "github.com/tolgaOzen/go-skeleton/pkg/database/postgres"

	"github.com/tolgaOzen/go-skeleton/internal"
)

// DataWriter - Structure for Data Writer
type DataWriter struct {
	database *db.Postgres
	// options
	txOptions pgx.TxOptions
}

func NewDataWriter(database *db.Postgres) *DataWriter {
	return &DataWriter{
		database:  database,
		txOptions: pgx.TxOptions{IsoLevel: pgx.ReadCommitted, AccessMode: pgx.ReadWrite},
	}
}

func (w *DataWriter) Write(ctx context.Context, name string) (err error) {
	// Start a new trace span and end it when the function exits.
	ctx, span := internal.Tracer.Start(ctx, "data-writer.write")
	defer span.End()

	slog.DebugContext(ctx, "write user")

	tx, err := w.database.WritePool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	// Build the SQL query using Squirrel
	builder := w.database.Builder.
		Insert("users").
		Columns("name").
		Values(name)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	slog.DebugContext(ctx, "successfully written user to the database")
	return nil
}
