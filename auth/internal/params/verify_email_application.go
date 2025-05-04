package params

import "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"

type VerifyEmailApp struct {
	EmailId    int64
	SecretCode string
}

type VerifyEmailAppResult struct {
	User        domain.User
	VerifyEmail domain.VerifyEmail
}
