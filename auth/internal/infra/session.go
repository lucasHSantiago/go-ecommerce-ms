package infra

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
	"github.com/rs/zerolog/log"
)

type SessionRepository struct {
	connPool DBTX
}

func NewSessionRepository(connPool DBTX) *SessionRepository {
	return &SessionRepository{connPool}
}

func getSessionError(err error, defaultReturn error, msg string) error {
	if errors.Is(err, ErrRecordNotFound) {
		return domain.ErrSessionNotFound
	}

	if pgError := GetPgError(err); pgError != nil {
		if pgError.Code == ForeignKeyViolation {
			switch pgError.ConstraintName {
			case "sessions_username_fkey":
				return domain.ErrUserNotFound
			}
		}
	}

	log.Error().Err(err).Msg(msg)
	return defaultReturn
}

const createSession = `
INSERT INTO sessions (
  id,
  username,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
`

func (s *SessionRepository) CreateSession(ctx context.Context, arg params.CreateSessionRepo) (*domain.Session, error) {
	args := []any{
		arg.ID,
		arg.Username,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
	}

	rows, _ := s.connPool.Query(ctx, createSession, args...)

	session, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.Session])
	if err != nil {
		return nil, getSessionError(err, domain.ErrCreateSession, "failed to create session")
	}

	return session, err
}

const getSession = `
SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
FROM sessions
WHERE id = $1 LIMIT 1
`

func (s *SessionRepository) GetSession(ctx context.Context, id uuid.UUID) (*domain.Session, error) {
	rows, _ := s.connPool.Query(ctx, getSession, id)

	session, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.Session])
	if err != nil {
		return nil, getSessionError(err, domain.ErrReadSession, "failed to get session")
	}

	return session, err
}
