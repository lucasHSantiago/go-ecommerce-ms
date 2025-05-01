package application

import (
	"context"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
)

type VerifyEmailRepository interface {
	CreateVerifyEmail(ctx context.Context, arg params.CreateVerifyEmailRepo) (*domain.VerifyEmail, error)
	UpdateVerifyEmail(ctx context.Context, arg params.UpdateVerifyEmailRepo) (*domain.VerifyEmail, error)
}
