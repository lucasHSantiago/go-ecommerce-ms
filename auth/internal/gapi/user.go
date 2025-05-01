package gapi

import (
	"context"
	"errors"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	user, err := server.userApplication.Create(ctx, toUserAppParams(req))
	if err != nil {
		var valErr *domain.ValidationErrors
		if errors.As(err, &valErr) {
			return nil, invalidArgumentError(*valErr)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	return toUserReqResponse(user), nil
}
