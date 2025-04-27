package application

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/handler/service"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/worker"
)

type UserApplication struct {
	userRepository  port.UserRepository
	taskDistributor worker.TaskDistributor
}

func NewUserApplication(userRepository port.UserRepository, taskDistributor worker.TaskDistributor) *UserApplication {
	return &UserApplication{userRepository, taskDistributor}
}

func (u *UserApplication) CreateUser(ctx context.Context, arg service.CreateUserParams) (*domain.User, error) {
	errValidation := validateCreateUserParams(arg)
	if errValidation != nil {
		return nil, fmt.Errorf("validation error: %w", errValidation)
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

func (u *UserApplication) UpdateUser(ctx context.Context, arg service.UpdateUserParams) (*domain.User, error) {
	errValidation := validateUpdateUserParams(arg)
	if errValidation != nil {
		return nil, fmt.Errorf("validation error: %w", errValidation)
	}

	updateUserParams := port.UpdateUserParams{}
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
