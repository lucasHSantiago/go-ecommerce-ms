package gapi

import (
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toUserResponse(user *domain.User) *gen.User {
	return &gen.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

func toCreateUserApp(req *gen.CreateUserRequest) params.CreateUserApp {
	return params.CreateUserApp{
		Username: req.GetUsername(),
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func toCreateUserResponse(user *domain.User) *gen.CreateUserResponse {
	return &gen.CreateUserResponse{
		User: toUserResponse(user),
	}
}

func toUpdateUserApp(req *gen.UpdateUserRequest) params.UpdateUserApp {
	return params.UpdateUserApp{
		Username: req.GetUsername(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}
}

func toUpdateUserResponse(user *domain.User) *gen.UpdateUserResponse {
	return &gen.UpdateUserResponse{
		User: toUserResponse(user),
	}
}
