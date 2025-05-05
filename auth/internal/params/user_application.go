package params

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateUserApp struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserApp struct {
	Username string  `json:"username"`
	FullName *string `json:"full_name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type LoginUserApp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserAppResult struct {
	User                  *domain.User `json:"user"`
	SessionId             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	RefreshToken          string       `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
}
