package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/mail"
	"github.com/rs/zerolog/log"
)

type TaskProcessor interface {
	Start() error
	Shutdown()
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server                *asynq.Server
	userRepository        application.UserRepository
	verifyEmailRepository application.VerifyEmailRepository
	mailer                mail.EmailSender
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, userRepository application.UserRepository, verifyEmailRepository application.VerifyEmailRepository, mailer mail.EmailSender) *RedisTaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				application.CriticalQueue: 10,
				application.DefaultQueue:  5,
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
