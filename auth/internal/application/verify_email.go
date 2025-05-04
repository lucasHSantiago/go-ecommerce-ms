package application

import (
	"context"
	"fmt"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
)

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

func validateVerifyEmailRequest(arg params.VerifyEmailApp) *domain.ValidationErrors {
	var errs domain.ValidationErrors

	if err := validateEmailId(arg.EmailId); err != nil {
		errs = append(errs, domain.NewFieldValidation("email_id", err))
	}

	if err := validateSecretCode(arg.SecretCode); err != nil {
		errs = append(errs, domain.NewFieldValidation("secret_code", err))
	}

	if len(errs) > 0 {
		return &errs
	}

	return nil
}

func validateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}

	return nil
}

func validateSecretCode(value string) error {
	return domain.ValidateString(value, 32, 128)
}
