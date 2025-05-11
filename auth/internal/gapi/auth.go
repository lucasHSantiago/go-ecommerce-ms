package gapi

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserApplication interface {
	Create(ctx context.Context, arg application.CreateUser) (*domain.User, error)
	Update(ctx context.Context, arg application.UpdateUser) (*domain.User, error)
	Login(ctx context.Context, arg application.LoginUser) (*application.LoginUserResult, error)
}

type VerifyEmailApplication interface {
	VerifyEmail(ctx context.Context, arg application.VerifyEmail) (*application.VerifyEmailResult, error)
}

type AuthServer struct {
	gen.UnimplementedAuthServiceServer
	userApplication        UserApplication
	verifyEmailApplication VerifyEmailApplication
}

func NewAuthServer(userApplication UserApplication, verVerifyEmailApplication VerifyEmailApplication) *AuthServer {
	return &AuthServer{
		userApplication:        userApplication,
		verifyEmailApplication: verVerifyEmailApplication,
	}
}

func (server *AuthServer) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	user, err := server.userApplication.Create(ctx, toCreateUserApp(req))
	if err != nil {
		var valErr validation.Errors
		if errors.As(err, &valErr) && valErr != nil {
			fmt.Println("entrou")
			return nil, invalidArgumentError(valErr)
		}
		log.Error().Err(err).Msg("failed to create user")
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	return toCreateUserResponse(user), nil
}

func (server *AuthServer) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error) {
	user, err := server.userApplication.Update(ctx, toUpdateUserApp(req))
	if err != nil {
		var valErr validation.Errors
		if errors.As(err, &valErr) && valErr != nil {
			return nil, invalidArgumentError(valErr)
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
		var valErr validation.Errors
		if errors.As(err, &valErr) && valErr != nil {
			return nil, invalidArgumentError(valErr)
		}
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		log.Error().Err(err).Msg("failed to login user")
		return nil, status.Errorf(codes.Internal, "failed to login user: %s", err)
	}

	return toLoginUserResponse(res), nil
}

func (server *AuthServer) VerifyEmail(ctx context.Context, req *gen.VerifyEmailRequest) (*gen.VerifyEmailResponse, error) {
	res, err := server.verifyEmailApplication.VerifyEmail(ctx, toVerifyEmailApp(req))
	if err != nil {
		var valErr validation.Errors
		if errors.As(err, &valErr) && valErr != nil {
			return nil, invalidArgumentError(valErr)
		}
		log.Error().Err(err).Msg("failed to verify email")
		return nil, status.Errorf(codes.Internal, "failed to verify email: %s", err)
	}

	return toVerifyEmailResponse(res), nil
}
