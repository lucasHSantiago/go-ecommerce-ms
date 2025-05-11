package application

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/infra"
)

type VerifyEmailRepository interface {
	CreateVerifyEmail(ctx context.Context, arg infra.CreateVerifyEmail) (*domain.VerifyEmail, error)
	UpdateVerifyEmail(ctx context.Context, arg infra.UpdateVerifyEmail) (*domain.VerifyEmail, error)
	VerifyEmailTx(ctx context.Context, arg infra.VerifyEmailTx) (infra.VerifyEmailTxResult, error)
}

type VerifyEmailApplication struct {
	verifyEmailRepository VerifyEmailRepository
}

func NewVerifyEmailApplication(verifyEmailRepository VerifyEmailRepository) *VerifyEmailApplication {
	return &VerifyEmailApplication{
		verifyEmailRepository: verifyEmailRepository,
	}
}

type VerifyEmail struct {
	EmailId    int64  `json:"email_id"`
	SecretCode string `json:"secret_code"`
}

type VerifyEmailResult struct {
	User        domain.User        `json:"user"`
	VerifyEmail domain.VerifyEmail `json:"verify_email"`
}

func (v *VerifyEmailApplication) VerifyEmail(ctx context.Context, arg VerifyEmail) (*VerifyEmailResult, error) {
	if errValidation := validateVerifyEmailRequest(arg); errValidation != nil {
		return nil, errValidation
	}

	txResult, err := v.verifyEmailRepository.VerifyEmailTx(ctx, infra.VerifyEmailTx{
		EmailId:     arg.EmailId,
		SecreteCode: arg.SecretCode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	response := &VerifyEmailResult{
		User:        *txResult.User,
		VerifyEmail: *txResult.VerifyEmail,
	}

	return response, nil
}

func validateVerifyEmailRequest(arg VerifyEmail) error {
	return validation.ValidateStruct(&arg,
		validation.Field(&arg.EmailId, validation.Required),
		validation.Field(&arg.SecretCode, validation.Required, validation.Length(32, 128)))
}
