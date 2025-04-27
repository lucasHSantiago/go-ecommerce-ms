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

type UserApplication interface {
	CreateUser(ctx context.Context, user CreateUserParams) (*domain.User, error)
}
