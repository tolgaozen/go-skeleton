package circuitBreaker

import (
	"context"

	"github.com/sony/gobreaker"

	"github.com/tolgaOzen/go-skeleton/internal/storage"
	"github.com/tolgaOzen/go-skeleton/pkg/database"
	base "github.com/tolgaOzen/go-skeleton/pkg/pb/base/v1"
)

// DataReader - Add circuit breaker behaviour to data reader
type DataReader struct {
	delegate storage.DataReader
	cb       *gobreaker.CircuitBreaker
}

// NewDataReader - Add circuit breaker behaviour to new data reader
func NewDataReader(delegate storage.DataReader, cb *gobreaker.CircuitBreaker) *DataReader {
	return &DataReader{delegate: delegate, cb: cb}
}

// ReadUsers - Read users with circuit breaker
func (r *DataReader) ReadUsers(ctx context.Context, pagination database.Pagination) ([]*base.User, error) {
	response, err := r.cb.Execute(func() (interface{}, error) {
		return r.delegate.ReadUsers(ctx, pagination)
	})
	if err != nil {
		return nil, err
	}
	return response.([]*base.User), nil
}
