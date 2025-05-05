package params

import "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"

type VerifyEmailApp struct {
	EmailId    int64  `json:"email_id"`
	SecretCode string `json:"secret_code"`
}

type VerifyEmailAppResult struct {
	User        domain.User        `json:"user"`
	VerifyEmail domain.VerifyEmail `json:"verify_email"`
}
