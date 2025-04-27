package application

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
)

type UserApplication struct {
	userRepository UserRepository
}

func NewUserApplication(userRepository UserRepository) *UserApplication {
	return &UserApplication{userRepository}
}

func (u *UserApplication) CreateUser(ctx context.Context, user domain.User) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err
	}

	return nil
	// hashedPassword, err := util.HashPassword(req.GetPassword())
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	// }
	//
	// arg := db.CreateUserTxParams{
	// 	CreateUserParams: db.CreateUserParams{
	// 		Username:       req.GetUsername(),
	// 		HashedPassword: hashedPassword,
	// 		FullName:       req.GetFullName(),
	// 		Email:          req.GetEmail(),
	// 	},
	// 	AfterCreate: func(user db.User) error {
	// 		taskPayload := &worker.PayloadSendVerifyEmail{
	// 			Username: user.Username,
	// 		}
	//
	// 		opts := []asynq.Option{
	// 			asynq.MaxRetry(10),
	// 			asynq.ProcessIn(10 * time.Second),
	// 			asynq.Queue(worker.CriticalQueue),
	// 		}
	//
	// 		return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
	// 	},
	// }
	//
	// txResult, err := server.store.CreateUserTx(ctx, arg)
	// if err != nil {
	// 	if db.ErrorCode(err) == db.UniqueViolation {
	// 		return nil, status.Errorf(codes.AlreadyExists, "%s", err)
	// 	}
	//
	// 	return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	// }
	//
	// response := &gen.CreateUserResponse{
	// 	User: convertUser(txResult.User),
	// }
	//
	// return response, nil
}
