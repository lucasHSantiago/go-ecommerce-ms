package application

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/distributor"
)

const (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
)

type PayloadSendVerifyEmail struct {
	Username string
}

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(ctx context.Context, payload *distributor.PayloadSendVerifyEmail, opts ...asynq.Option) error
}
