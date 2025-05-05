package params

import (
	"time"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateUserRepo struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

type UpdateUserRepo struct {
	HashedPassword    *string    `json:"hashed_password"`
	PasswordChangedAt *time.Time `json:"password_changed_at"`
	FullName          *string    `json:"full_name"`
	Email             *string    `json:"email"`
	IsEmailVerified   *bool      `json:"is_email_verified"`
	Username          string     `json:"username"`
}

type CreateUserTxRepo struct {
	CreateUserRepo `json:"create_user_repo"`
	AfterCreate    func(user domain.User) error `json:"after_create"`
}

type CreateUserTxRepoResult struct {
	User domain.User `json:"user"`
}
