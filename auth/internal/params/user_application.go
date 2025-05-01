package params

import (
	"time"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateUserApp struct {
	Username string
	FullName string
	Email    string
	Password string
}

type UpdateUserApp struct {
	Username string
	FullName *string
	Email    *string
	Password *string
}

type LoginUserApp struct {
	Username string
	Password string
}

type LoginUserAppResult struct {
	User                  *domain.User
	SessionId             uuid.UUID
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  *time.Time
	RefreshTokenExpiresAt *time.Time
}
