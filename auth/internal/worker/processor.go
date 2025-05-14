package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/infra"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/mail"
	"github.com/rs/zerolog/log"
)

const (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
)

type UserReader interface {
	GetUser(ctx context.Context, username string) (*domain.User, error)
}

type VerifyEmailCreator interface {
	CreateVerifyEmail(ctx context.Context, arg infra.CreateVerifyEmail) (*domain.VerifyEmail, error)
}

type RedisTaskProcessor struct {
	server                *asynq.Server
	userRepository        UserReader
	verifyEmailRepository VerifyEmailCreator
	mailer                mail.EmailSender
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, userRepository UserReader, verifyEmailRepository VerifyEmailCreator, mailer mail.EmailSender) *RedisTaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				CriticalQueue: 10,
				DefaultQueue:  5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(_ context.Context, task *asynq.Task, err error) {
				log.Error().
					Err(err).
					Str("type", task.Type()).
					Bytes("payload", task.Payload()).
					Msg("process task failed")
			}),
			Logger: NewLogger(),
		},
	)

	return &RedisTaskProcessor{
		server:                server,
		userRepository:        userRepository,
		verifyEmailRepository: verifyEmailRepository,
		mailer:                mailer,
	}
}

func (r *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, r.ProcessTaskSendVerifyEmail)

	return r.server.Start(mux)
}

func (r *RedisTaskProcessor) Shutdown() {
	r.server.Shutdown()
}
