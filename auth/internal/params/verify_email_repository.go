package params

import "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"

type CreateVerifyEmailRepo struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	SecretCode string `json:"secret_code"`
}

type UpdateVerifyEmailRepo struct {
	ID         int64  `json:"id"`
	SecretCode string `json:"secret_code"`
}

type VerifyEmailTxRepo struct {
	EmailId     int64  `json:"email_id"`
	SecreteCode string `json:"secrete_code"`
}

type VerifyEmailTxRepoResult struct {
	User        *domain.User        `json:"user"`
	VerifyEmail *domain.VerifyEmail `json:"verify_email"`
}
