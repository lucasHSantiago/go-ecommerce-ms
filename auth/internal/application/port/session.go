package port

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type CreateSessionParams struct {
	ID           uuid.UUID
	Username     string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
}

type SessionRepository interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (*domain.Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (*domain.Session, error)
}
