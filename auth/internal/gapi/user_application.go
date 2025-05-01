package gapi

import (
	"context"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
)

type UserApplication interface {
	Create(ctx context.Context, arg params.CreateUserApp) (*domain.User, error)
	Update(ctx context.Context, arg params.UpdateUserApp) (*domain.User, error)
	Login(ctx context.Context, arg params.LoginUserApp) (*params.LoginUserAppResult, error)
}
