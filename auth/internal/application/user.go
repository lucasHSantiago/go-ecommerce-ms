package application

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/gapi/service"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/token"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/worker"
)

type UserApplication struct {
	userRepository     port.UserRepository
	sessionRespository port.SessionRepository
	taskDistributor    worker.TaskDistributor
	tokenMaker         token.Maker
	config             *util.Config
}

func NewUserApplication(userRepository port.UserRepository, sessionRepository port.SessionRepository, taskDistributor worker.TaskDistributor, tokenMaker token.Maker, config *util.Config) service.UserApplication {
	return &UserApplication{
		userRepository:     userRepository,
		sessionRespository: sessionRepository,
		taskDistributor:    taskDistributor,
		tokenMaker:         tokenMaker,
		config:             config,
	}
}

func (u *UserApplication) Create(ctx context.Context, arg service.CreateUserParams) (*domain.User, error) {
	if errValidation := validateCreateUserParams(arg); errValidation != nil {
		return nil, errValidation
	}

	hashedPassword, err := util.HashPassword(arg.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	argTx := port.CreateUserTxParams{
		CreateUserParams: port.CreateUserParams{
			Username:       arg.Username,
			HashedPassword: hashedPassword,
			FullName:       arg.FullName,
			Email:          arg.Email,
		},
		AfterCreate: func(user domain.User) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				Username: user.Username,
			}

			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.CriticalQueue),
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

func (u *UserApplication) Update(ctx context.Context, arg service.UpdateUserParams) (*domain.User, error) {
	if errValidation := validateUpdateUserParams(arg); errValidation != nil {
		return nil, errValidation
	}

	updateUserParams := port.UpdateUserParams{
		FullName: arg.FullName,
		Username: arg.Username,
		Email:    arg.Email,
	}

	if arg.Password != nil {
		hashedPassword, err := util.HashPassword(*arg.Password)
		if err != nil {
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

// @TODO: testar
func (u *UserApplication) Login(ctx context.Context, arg service.LoginUserParams) (*service.LoginUserResult, error) {
	if errValidation := validateLoginUserParams(arg); errValidation != nil {
		return nil, fmt.Errorf("validation error: %w", errValidation)
	}

	user, err := u.userRepository.GetUser(ctx, arg.Username)
	if err != nil {
		return nil, fmt.Errorf("username invalid: %w", err)
	}

	err = util.CheckPassword(arg.Password, user.HashedPassword)
	if err != nil {
		return nil, fmt.Errorf("invalid usarname or password")
	}

	accessToken, accessPayload, err := u.tokenMaker.CreateToken(user.Username, user.Role, u.config.AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %w", err)
	}

	refreshToken, refreshPayload, err := u.tokenMaker.CreateToken(user.Username, user.Role, u.config.RefreshTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: %w", err)
	}

	metadata := util.ExtractMetadata(ctx)
	session, err := u.sessionRespository.CreateSession(ctx, port.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    metadata.UserAgent,
		ClientIp:     metadata.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	response := &service.LoginUserResult{
		User:                  user,
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  &accessPayload.ExpiredAt,
		RefreshTokenExpiresAt: &refreshPayload.ExpiredAt,
	}

	return response, nil
}
