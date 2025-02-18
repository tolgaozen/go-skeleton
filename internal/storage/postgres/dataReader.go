package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"

	"github.com/tolgaOzen/go-skeleton/pkg/database"
	db "github.com/tolgaOzen/go-skeleton/pkg/database/postgres"

	"github.com/tolgaOzen/go-skeleton/internal"
	"github.com/tolgaOzen/go-skeleton/internal/storage"
	basev1 "github.com/tolgaOzen/go-skeleton/pkg/pb/base/v1"
)

// DataReader is a struct which holds a reference to the database, transaction options and a logger.
// It is responsible for reading data from the database.
type DataReader struct {
	database  *db.Postgres  // database is an instance of the PostgreSQL database
	txOptions pgx.TxOptions // txOptions specifies the isolation level for database transaction and sets it as read only
}

// NewDataReader is a constructor function for DataReader.
// It initializes a new DataReader with a given database, a logger, and sets transaction options to be read-only with Repeatable Read isolation level.
func NewDataReader(database *db.Postgres) *DataReader {
	return &DataReader{
		database:  database,                                                             // Set the database to the passed in PostgreSQL instance
		txOptions: pgx.TxOptions{IsoLevel: pgx.ReadCommitted, AccessMode: pgx.ReadOnly}, // Set the transaction options
	}
}

// ReadUsers reads relation tuples from the storage based on the given filter.
func (r *DataReader) ReadUsers(ctx context.Context, pagination database.Pagination) (users []*basev1.User, err error) {
	// Start a new trace span and end it when the function exits.
	ctx, span := internal.Tracer.Start(ctx, "data-reader.read-users")
	defer span.End()

	slog.DebugContext(ctx, "querying users")

	var args []interface{}
	builder := r.database.Builder.
		Select("id, name, created_at").
		From(UsersTable).
		OrderBy("created_at DESC").
		Limit(uint64(pagination.Size())).
		Offset(uint64(pagination.Size()) * uint64(max(0, pagination.Page()-1)))

	// Generate the SQL query and arguments.
	var query string
	query, args, err = builder.ToSql()
	if err != nil {
		return nil, err
	}

	slog.DebugContext(ctx, "generated sql query", slog.String("query", query), slog.Any("arguments", args))

	// Execute the SQL query and retrieve the result rows.
	var rows pgx.Rows
	rows, err = r.database.ReadPool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		fnd := &storage.User{}
		err = rows.Scan(
			&fnd.ID,
			&fnd.Name,
			&fnd.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, fnd.ToProto())
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	slog.DebugContext(ctx, "successfully retrieved and converted users from the database")

	// Return the results.
	return users, nil
}
