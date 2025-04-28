package infra

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/rs/zerolog/log"
)

type UserRepository struct {
	connPool DBTX
}

func NewUserRepository(connPool DBTX) *UserRepository {
	return &UserRepository{connPool}
}

func getUserError(err error, msg string) error {
	if errors.Is(err, ErrRecordNotFound) {
		return domain.ErrUserNotFound
	}

	if pgError := GetPgError(err); pgError != nil {
		if pgError.Code == UniqueViolation {
			switch pgError.ConstraintName {
			case "users_pkey":
				return domain.ErrUsernameAlreadyExist
			case "users_email_key":
				return domain.ErrEmailAlreadyExist
			}
		}
	}

	log.Error().Err(err).Msg(msg)
	return err
}

const createUser = `
INSERT INTO users (
	username,
	hashed_password,
	full_name,
email
) VALUES (
	$1, $2, $3, $4
) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at, is_email_verified, role
`

func (u *UserRepository) CreateUser(ctx context.Context, arg port.CreateUserParams) (*domain.User, error) {
	args := []any{
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	}

	rows, _ := u.connPool.Query(ctx, createUser, args...)

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])
	if err != nil {
		return nil, getUserError(err, "faield to create user")
	}

	return user, err
}

const getUser = `
SELECT username, hashed_password, full_name, email, password_changed_at, created_at, is_email_verified, role FROM users
WHERE username = $1 LIMIT 1
`

func (u *UserRepository) GetUser(ctx context.Context, username string) (*domain.User, error) {
	rows, _ := u.connPool.Query(ctx, getUser, username)

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])
	if err != nil {
		return nil, getUserError(err, "failed to get user by username")
	}

	return user, err
}

const updateUser = `
UPDATE users
SET
    hashed_password = COALESCE($1, hashed_password),
    password_changed_at = COALESCE($2, password_changed_at),
    full_name = COALESCE($3, full_name),
    email = COALESCE($4, email),
    is_email_verified = COALESCE($5, is_email_verified)
WHERE
    username = $6
RETURNING username, hashed_password, full_name, email, password_changed_at, created_at, is_email_verified, role
`

func (u *UserRepository) UpdateUser(ctx context.Context, arg port.UpdateUserParams) (*domain.User, error) {
	args := []any{
		util.StringToText(arg.HashedPassword),
		util.TimeToTimestamptz(arg.PasswordChangedAt),
		util.StringToText(arg.FullName),
		util.StringToText(arg.Email),
		util.BoolToBool(arg.IsEmailVerified),
		arg.Username,
	}

	rows, _ := u.connPool.Query(ctx, updateUser, args...)

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[domain.User])
	if err != nil {
		return nil, getUserError(err, "failed to update user")
	}

	return user, err
}

func (u *UserRepository) CreateUserTx(ctx context.Context, arg port.CreateUserTxParams) (port.CreateUserTxResult, error) {
	var result port.CreateUserTxResult

	err := execTx(ctx, u.connPool, func(tx pgx.Tx) error {
		userRepository := NewUserRepository(tx)
		user, err := userRepository.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		result.User = *user

		return arg.AfterCreate(result.User)
	})

	return result, err
}
