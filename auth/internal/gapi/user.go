package gapi

import (
	"context"
	"errors"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	user, err := server.userApplication.Create(ctx, toCreateUserApp(req))
	if err != nil {
		var valErr *domain.ValidationErrors
		if errors.As(err, &valErr) {
			return nil, invalidArgumentError(*valErr)
		}
		log.Error().Err(err).Msg("failed to create user")
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	return toCreateUserResponse(user), nil
}

func (server *Server) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error) {
	user, err := server.userApplication.Update(ctx, toUpdateUserApp(req))
	if err != nil {
		var valErr *domain.ValidationErrors
		if errors.As(err, &valErr) {
			return nil, invalidArgumentError(*valErr)
		}
		log.Error().Err(err).Msg("failed to update user")
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	return toUpdateUserResponse(user), nil
}
