package service

import (
	"context"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateUserParams struct {
	Username string
	Password string
	FullName string
	Email    string
}

type UpdateUserParams struct {
	Username string
	FullName *string
	Email    *string
	Password *string
}

type UserApplication interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (*domain.User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*domain.User, error)
}
