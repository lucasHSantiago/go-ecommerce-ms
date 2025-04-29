package infra

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/rs/zerolog/log"
)

type VerifyEmailRepository struct {
	connPool DBTX
}

func NewVerifyEmailRepository(connPool DBTX) port.VerifyEmailRepository {
	return &VerifyEmailRepository{connPool}
}

func getVerifyEmailError(err error, msg string) error {
	if pgError := GetPgError(err); pgError != nil {
		if pgError.Code == ForeignKeyViolation {
			switch pgError.ConstraintName {
			case "verify_emails_username_fkey":
				return domain.ErrUserNotFound
			}
		}
	}

	log.Error().Err(err).Msg(msg)
	return err
}

const createVerifyEmail = `
INSERT INTO "verify_emails" (
    username,
    email,
    secret_code
) VALUES (
    $1, $2, $3
) RETURNING id, username, email, secret_code, is_used, created_at, expired_at
`

func (v *VerifyEmailRepository) CreateVerifyEmail(ctx context.Context, arg port.CreateVerifyEmailParams) (*domain.VerifyEmail, error) {
	args := []any{
		arg.Username,
		arg.Email,
		arg.SecretCode,
	}

	rows, _ := v.connPool.Query(ctx, createVerifyEmail, args...)

	verifyEmail, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.VerifyEmail])
	if err != nil {
		return nil, getVerifyEmailError(err, "faield to create user")
	}

	return verifyEmail, err
}

const updateVerifyEmail = `-- name: UpdateVerifyEmail :one
UPDATE verify_emails 
SET is_used = true
WHERE id = $1
AND secret_code = $2
AND is_used = FALSE
AND expired_at > now()
RETURNING id, username, email, secret_code, is_used, created_at, expired_at
`

func (v *VerifyEmailRepository) UpdateVerifyEmail(ctx context.Context, arg port.UpdateVerifyEmailParams) (*domain.VerifyEmail, error) {
	args := []any{
		arg.ID,
		arg.SecretCode,
	}

	rows, _ := v.connPool.Query(ctx, updateVerifyEmail, args...)

	verifyEmail, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.VerifyEmail])
	if err != nil {
		return nil, getVerifyEmailError(err, "failed to update verify email")
	}

	return verifyEmail, nil
}
