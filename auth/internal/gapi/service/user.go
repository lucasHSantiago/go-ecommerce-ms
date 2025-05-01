package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateUserParams struct {
	Username string
	FullName string
	Email    string
	Password string
}

type UpdateUserParams struct {
	Username string
	FullName *string
	Email    *string
	Password *string
}

type LoginUserParams struct {
	Username string
	Password string
}

type LoginUserResult struct {
	User                  *domain.User
	SessionId             uuid.UUID
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  *time.Time
	RefreshTokenExpiresAt *time.Time
}

type UserApplication interface {
	Create(ctx context.Context, arg CreateUserParams) (*domain.User, error)
	Update(ctx context.Context, arg UpdateUserParams) (*domain.User, error)
	Login(ctx context.Context, arg LoginUserParams) (*LoginUserResult, error)
}
