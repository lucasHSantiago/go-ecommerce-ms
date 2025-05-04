package gapi

import (
	"context"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
)

type VerifyEmailApplication interface {
	VerifyEmail(ctx context.Context, arg params.VerifyEmailApp) (*params.VerifyEmailAppResult, error)
}
