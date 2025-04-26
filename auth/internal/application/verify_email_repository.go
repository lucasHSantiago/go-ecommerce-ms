package application

import (
	"context"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateVerifyEmailParams struct {
	Username   string
	Email      string
	SecretCode string
}

type UpdateVerifyEmailParams struct {
	ID         int64
	SecretCode string
}

type VerifyEmailRepository interface {
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (*domain.VerifyEmail, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (*domain.VerifyEmail, error)
}
