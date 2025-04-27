package port

import (
	"context"
	"time"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateUserParams struct {
	Username       string
	HashedPassword string
	FullName       string
	Email          string
}

type UpdateUserParams struct {
	HashedPassword    *string
	PasswordChangedAt *time.Time
	FullName          *string
	Email             *string
	IsEmailVerified   *bool
	Username          string
}

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user domain.User) error
}

type CreateUserTxResult struct {
	User domain.User
}

type UserRepository interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (*domain.User, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	GetUser(ctx context.Context, username string) (*domain.User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*domain.User, error)
}
