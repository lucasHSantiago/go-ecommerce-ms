package params

import (
	"time"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateUserRepo struct {
	Username       string
	HashedPassword string
	FullName       string
	Email          string
}

type UpdateUserRepo struct {
	HashedPassword    *string
	PasswordChangedAt *time.Time
	FullName          *string
	Email             *string
	IsEmailVerified   *bool
	Username          string
}

type CreateUserTxRepo struct {
	CreateUserRepo
	AfterCreate func(user domain.User) error
}

type CreateUserTxRepoResult struct {
	User domain.User
}
