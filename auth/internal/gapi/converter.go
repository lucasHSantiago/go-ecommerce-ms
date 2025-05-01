package gapi

import (
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toUserAppParams(req *gen.CreateUserRequest) params.CreateUserApp {
	return params.CreateUserApp{
		Username: req.GetUsername(),
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func toUserReqResponse(user *domain.User) *gen.CreateUserResponse {
	return &gen.CreateUserResponse{
		User: &gen.User{
			Username:          user.Username,
			FullName:          user.FullName,
			Email:             user.Email,
			PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
			CreatedAt:         timestamppb.New(user.CreatedAt),
		},
	}
}
