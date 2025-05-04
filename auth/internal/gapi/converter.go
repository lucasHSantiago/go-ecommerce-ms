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

func toLoginUserApp(req *gen.LoginUserRequest) params.LoginUserApp {
	return params.LoginUserApp{
		Username: req.Username,
		Password: req.Password,
	}
}

func toLoginUserResponse(res *params.LoginUserAppResult) *gen.LoginUserResponse {
	return &gen.LoginUserResponse{
		User:                  toUserResponse(res.User),
		SessionId:             res.SessionId.String(),
		AccessToken:           res.AccessToken,
		RefreshToken:          res.RefreshToken,
		AccessTokenExpiresAt:  timestamppb.New(res.AccessTokenExpiresAt),
		RefreshTokenExpiresAt: timestamppb.New(res.RefreshTokenExpiresAt),
	}
}

func toVerifyEmailApp(req *gen.VerifyEmailRequest) params.VerifyEmailApp {
	return params.VerifyEmailApp{
		EmailId:    req.GetEmailId(),
		SecretCode: req.GetSecretCode(),
	}
}

func toVerifyEmailResponse(res *params.VerifyEmailAppResult) *gen.VerifyEmailResponse {
	return &gen.VerifyEmailResponse{
		IsVerified: res.User.IsEmailVerified,
	}
}
