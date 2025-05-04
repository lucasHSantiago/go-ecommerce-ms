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

type AuthServer struct {
	gen.UnimplementedAuthServiceServer
	userApplication UserApplication
}

func NewAuthServer(userApplication UserApplication) *AuthServer {
	return &AuthServer{
		userApplication: userApplication,
	}
}

func (server *AuthServer) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
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

func (server *AuthServer) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error) {
	user, err := server.userApplication.Update(ctx, toUpdateUserApp(req))
	if err != nil {
		var valErr *domain.ValidationErrors
		if errors.As(err, &valErr) {
			return nil, invalidArgumentError(*valErr)
		}
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		log.Error().Err(err).Msg("failed to update user")
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	return toUpdateUserResponse(user), nil
}

func (server *AuthServer) LoginUser(ctx context.Context, req *gen.LoginUserRequest) (*gen.LoginUserResponse, error) {
	res, err := server.userApplication.Login(ctx, toLoginUserApp(req))
	if err != nil {
		var valErr *domain.ValidationErrors
		if errors.As(err, &valErr) {
			return nil, invalidArgumentError(*valErr)
		}
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		log.Error().Err(err).Msg("failed to login user")
		return nil, status.Errorf(codes.Internal, "failed to login user: %s", err)
	}

	return toLoginUserResponse(res), nil
}
