package infra

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/rs/zerolog/log"
)

type VerifyEmailRepository struct {
	connPool DBTX
}

func NewVerifyEmailRepository(connPool DBTX) *VerifyEmailRepository {
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

type CreateVerifyEmail struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
}

func (v *VerifyEmailRepository) CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmail) (*domain.VerifyEmail, error) {
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

type UpdateVerifyEmail struct {
	ID         int64  `json:"id"`
	SecretCode string `json:"secret_code"`
}

func (v *VerifyEmailRepository) UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmail) (*domain.VerifyEmail, error) {
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

type VerifyEmailTx struct {
	EmailId     int64  `json:"email_id"`
	SecreteCode string `json:"secrete_code"`
}

type VerifyEmailTxResult struct {
	User        *domain.User        `json:"user"`
	VerifyEmail *domain.VerifyEmail `json:"verify_email"`
}

func (v *VerifyEmailRepository) VerifyEmailTx(ctx context.Context, arg VerifyEmailTx) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := execTx(ctx, v.connPool, func(tx pgx.Tx) error {
		var err error

		verifyEmailRepository := NewVerifyEmailRepository(tx)
		result.VerifyEmail, err = verifyEmailRepository.UpdateVerifyEmail(ctx, UpdateVerifyEmail{
			ID:         arg.EmailId,
			SecretCode: arg.SecreteCode,
		})
		if err != nil {
			return err
		}

		isEmailVerified := true

		userRepository := NewUserRepository(tx)
		result.User, err = userRepository.UpdateUser(ctx, UpdateUser{
			Username:        result.VerifyEmail.Username,
			IsEmailVerified: &isEmailVerified,
		})

		return err
	})

	return result, err
}
