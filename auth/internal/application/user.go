package application

import (
	"context"
	"net/mail"
	"regexp"
	"time"

	"github.com/hibiken/asynq"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/handler/service"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/validator"
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
	errValidation := ValidateCreateUserParams(arg)
	if errValidation != nil {
		return nil, errValidation
	}

	hashedPassword, err := util.HashPassword(arg.Password)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &res.User, nil
}

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[A-Za-z ]+$`).MatchString
)

func ValidateCreateUserParams(arg service.CreateUserParams) *validator.ValidationErrors {
	var errs validator.ValidationErrors

	if len(arg.Username) < 3 || len(arg.Username) > 100 || !isValidUsername(arg.Username) {
		errs = append(errs, validator.NewFielValidation("username", "username must have 3-100 characters and contain only letters, numbers, or underscores"))
	}

	if len(arg.Password) < 6 || len(arg.Password) > 100 {
		errs = append(errs, validator.NewFielValidation("password", "password must have 6-100 characters"))
	}

	if len(arg.FullName) < 3 || len(arg.FullName) > 100 || !isValidFullName(arg.FullName) {
		errs = append(errs, validator.NewFielValidation("full name", "full name must have 3-100 characters and contain only letters and spaces"))
	}

	if _, err := mail.ParseAddress(arg.Email); err != nil {
		errs = append(errs, validator.NewFielValidation("email", "email must be a valid email address"))
	}

	if len(errs) > 0 {
		return &errs
	}

	return nil
}
