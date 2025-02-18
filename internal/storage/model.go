package storage

import (
	"time"

	basev1 "github.com/tolgaOzen/go-skeleton/pkg/pb/base/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// User is the model for the user entity.
type User struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
}

// ToProto - Convert database user to base user
func (r User) ToProto() *basev1.User {
	return &basev1.User{
		Id:        r.ID,
		Name:      r.Name,
		CreatedAt: timestamppb.New(r.CreatedAt),
	}
}
