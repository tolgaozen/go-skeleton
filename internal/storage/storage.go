package storage

import (
	"context"

	"github.com/tolgaOzen/go-skeleton/pkg/database"
	basev1 "github.com/tolgaOzen/go-skeleton/pkg/pb/base/v1"
)

// DataReader - Interface for reading Data from the storage.
type DataReader interface {
	// ReadUsers - Read users from the storage.
	ReadUsers(ctx context.Context, pagination database.Pagination) (users []*basev1.User, err error)
}

type NoopDataReader struct{}

func NewNoopDataReader() DataReader {
	return &NoopDataReader{}
}

func (f *NoopDataReader) ReadUsers(_ context.Context, _ database.Pagination) ([]*basev1.User, error) {
	return []*basev1.User{}, nil
}

type DataWriter interface {
	Write(ctx context.Context, name string) (err error)
}

type NoopDataWriter struct{}

func NewNoopDataWriter() DataWriter {
	return &NoopDataWriter{}
}

func (n *NoopDataWriter) Write(_ context.Context, _ string) error {
	return nil
}
