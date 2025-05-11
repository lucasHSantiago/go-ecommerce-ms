package application

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/distributor"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/infra"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/pkg/token"
	"github.com/rs/zerolog/log"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg infra.CreateUser) (*domain.User, error)
	CreateUserTx(ctx context.Context, arg infra.CreateUserTx) (infra.CreateUserTxResult, error)
	GetUser(ctx context.Context, username string) (*domain.User, error)
	UpdateUser(ctx context.Context, arg infra.UpdateUser) (*domain.User, error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context, arg infra.CreateSession) (*domain.Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (*domain.Session, error)
}

type JwtTokenMaker interface {
	CreateToken(username string, role string, duration time.Duration) (string, *token.Payload, error)
}

type UserApplication struct {
	userRepository     UserRepository
	sessionRespository SessionRepository
	taskDistributor    TaskDistributor
	tokenMaker         JwtTokenMaker
	config             *util.Config
}

func NewUserApplication(userRepository UserRepository, sessionRepository SessionRepository, taskDistributor TaskDistributor, tokenMaker JwtTokenMaker, config *util.Config) *UserApplication {
	return &UserApplication{
		userRepository:     userRepository,
		sessionRespository: sessionRepository,
		taskDistributor:    taskDistributor,
		tokenMaker:         tokenMaker,
		config:             config,
	}
}

type CreateUser struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserApplication) Create(ctx context.Context, arg CreateUser) (*domain.User, error) {
	if errValidation := validateCreateUserParams(arg); errValidation != nil {
		return nil, errValidation
	}

	hashedPassword, err := util.HashPassword(arg.Password)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	argTx := infra.CreateUserTx{
		CreateUser: infra.CreateUser{
			Username:       arg.Username,
			HashedPassword: hashedPassword,
			FullName:       arg.FullName,
			Email:          arg.Email,
		},
		AfterCreate: func(user domain.User) error {
			taskPayload := &distributor.PayloadSendVerifyEmail{
				Username: user.Username,
			}

			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(CriticalQueue),
			}

			return u.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}

	res, err := u.userRepository.CreateUserTx(ctx, argTx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &res.User, nil
}

type UpdateUser struct {
	Username string  `json:"username"`
	FullName *string `json:"full_name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (u *UserApplication) Update(ctx context.Context, arg UpdateUser) (*domain.User, error) {
	if errValidation := validateUpdateUserParams(arg); errValidation != nil {
		return nil, errValidation
	}

	updateUserParams := infra.UpdateUser{
		FullName: arg.FullName,
		Username: arg.Username,
		Email:    arg.Email,
	}

	if arg.Password != nil {
		hashedPassword, err := util.HashPassword(*arg.Password)
		if err != nil {
			log.Error().Err(err).Msg("failed to hash password")
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}

		updateUserParams.HashedPassword = &hashedPassword

		now := time.Now()
		updateUserParams.PasswordChangedAt = &now
	}

	user, err := u.userRepository.UpdateUser(ctx, updateUserParams)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResult struct {
	User                  *domain.User `json:"user"`
	SessionId             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	RefreshToken          string       `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
}

func (u *UserApplication) Login(ctx context.Context, arg LoginUser) (*LoginUserResult, error) {
	if errValidation := validateLoginUserParams(arg); errValidation != nil {
		return nil, errValidation
	}

	user, err := u.userRepository.GetUser(ctx, arg.Username)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(arg.Password, user.HashedPassword)
	if err != nil {
		return nil, ErrInvalidLoginPassword
	}

	accessToken, accessPayload, err := u.tokenMaker.CreateToken(user.Username, user.Role, u.config.AccessTokenDuration)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		return nil, fmt.Errorf("failed to create token: %w", err)
	}

	refreshToken, refreshPayload, err := u.tokenMaker.CreateToken(user.Username, user.Role, u.config.RefreshTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	metadata := util.ExtractMetadata(ctx)
	session, err := u.sessionRespository.CreateSession(ctx, infra.CreateSession{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    metadata.UserAgent,
		ClientIp:     metadata.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, err
	}

	response := &LoginUserResult{
		User:                  user,
		SessionId:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}

	return response, nil
}
