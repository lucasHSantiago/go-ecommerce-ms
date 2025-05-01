package application

import (
	"context"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg params.CreateUserRepo) (*domain.User, error)
	CreateUserTx(ctx context.Context, arg params.CreateUserTxRepo) (params.CreateUserTxRepoResult, error)
	GetUser(ctx context.Context, username string) (*domain.User, error)
	UpdateUser(ctx context.Context, arg params.UpdateUserRepo) (*domain.User, error)
}
