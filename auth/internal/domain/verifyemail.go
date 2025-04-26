package domain

import "time"

type VerifyEmail struct {
	ID         int64
	Username   string
	Email      string
	SecretCode string
	IsUsed     bool
	CreatedAt  time.Time
	ExpiredAt  time.Time
}
