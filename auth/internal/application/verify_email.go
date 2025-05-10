package application

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
)

type VerifyEmailRepository interface {
	CreateVerifyEmail(ctx context.Context, arg params.CreateVerifyEmailRepo) (*domain.VerifyEmail, error)
	UpdateVerifyEmail(ctx context.Context, arg params.UpdateVerifyEmailRepo) (*domain.VerifyEmail, error)
	VerifyEmailTx(ctx context.Context, arg params.VerifyEmailTxRepo) (params.VerifyEmailTxRepoResult, error)
}

type VerifyEmailApplication struct {
	verifyEmailRepository VerifyEmailRepository
}

func NewVerifyEmailApplication(verifyEmailRepository VerifyEmailRepository) *VerifyEmailApplication {
	return &VerifyEmailApplication{
		verifyEmailRepository: verifyEmailRepository,
	}
}

func (v *VerifyEmailApplication) VerifyEmail(ctx context.Context, arg params.VerifyEmailApp) (*params.VerifyEmailAppResult, error) {
	if errValidation := validateVerifyEmailRequest(arg); errValidation != nil {
		return nil, errValidation
	}

	txResult, err := v.verifyEmailRepository.VerifyEmailTx(ctx, params.VerifyEmailTxRepo{
		EmailId:     arg.EmailId,
		SecreteCode: arg.SecretCode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	response := &params.VerifyEmailAppResult{
		User:        *txResult.User,
		VerifyEmail: *txResult.VerifyEmail,
	}

	return response, nil
}

func validateVerifyEmailRequest(arg params.VerifyEmailApp) error {
	return validation.ValidateStruct(&arg,
		validation.Field(&arg.EmailId, validation.Required),
		validation.Field(&arg.SecretCode, validation.Required, validation.Length(32, 128)))
}
