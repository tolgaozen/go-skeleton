package servers

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/status"

	"github.com/tolgaOzen/go-skeleton/internal"
	"github.com/tolgaOzen/go-skeleton/internal/storage"
	"github.com/tolgaOzen/go-skeleton/pkg/database"
	v1 "github.com/tolgaOzen/go-skeleton/pkg/pb/base/v1"
)

// UserServer - Structure for User Server
type UserServer struct {
	v1.UnimplementedUserServiceServer

	dr storage.DataReader
	dw storage.DataWriter
}

// NewUserServer - Creates new User Server
func NewUserServer(dr storage.DataReader, dw storage.DataWriter) *UserServer {
	return &UserServer{
		dr: dr,
		dw: dw,
	}
}

// Create - Create new Tenant
func (t *UserServer) Create(ctx context.Context, request *v1.UserCreateRequest) (*v1.MessageResponse, error) {
	ctx, span := internal.Tracer.Start(ctx, "user.create")
	defer span.End()

	err := t.dw.Write(ctx, request.GetName())
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		slog.ErrorContext(ctx, err.Error())
		return nil, status.Error(GetStatus(err), err.Error())
	}

	return &v1.MessageResponse{
		Message: "success",
	}, nil
}

// List - List Users
func (t *UserServer) List(ctx context.Context, request *v1.UserListRequest) (*v1.UserListResponse, error) {
	ctx, span := internal.Tracer.Start(ctx, "user.list")
	defer span.End()

	users, err := t.dr.ReadUsers(ctx, database.NewPagination(database.Size(request.GetSize()), database.Page(request.GetPage())))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		slog.ErrorContext(ctx, err.Error())
		return nil, status.Error(GetStatus(err), err.Error())
	}

	return &v1.UserListResponse{
		Users: users,
	}, nil
}
