package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, arg params.CreateSessionRepo) (*domain.Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (*domain.Session, error)
}
