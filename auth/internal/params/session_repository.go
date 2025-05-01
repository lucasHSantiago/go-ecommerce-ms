package params

import (
	"time"

	"github.com/google/uuid"
)

type CreateSessionRepo struct {
	ID           uuid.UUID
	Username     string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
}
