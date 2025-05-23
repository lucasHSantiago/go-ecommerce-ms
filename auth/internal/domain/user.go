package domain

import (
	"errors"
	"time"
)

var (
	ErrUsernameAlreadyExist = errors.New("username already exists")
	ErrEmailAlreadyExist    = errors.New("email already exists")
	ErrUserNotFound         = errors.New("user not found")
	ErrReadUser             = errors.New("failed to get user")
	ErrCreateUser           = errors.New("failed to create user")
	ErrUpdateUser           = errors.New("failed to update user")
)

const (
	UserRole  = "user"
	AdminRole = "admin"
)

type User struct {
	Username          string    `db:"username" `
	Role              string    `db:"role"`
	HashedPassword    string    `db:"hashed_password" `
	FullName          string    `db:"full_name" `
	Email             string    `db:"email" `
	IsEmailVerified   bool      `db:"is_email_verified"`
	PasswordChangedAt time.Time `db:"password_changed_at"`
	CreatedAt         time.Time `db:"created_at"`
}
