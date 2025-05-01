package distributor

import (
	"context"

	"github.com/hibiken/asynq"
)

type PayloadSendVerifyEmail struct {
	Username string
}

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opts ...asynq.Option) error
}
