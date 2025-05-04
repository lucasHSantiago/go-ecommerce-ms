package params

import "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"

type CreateVerifyEmailRepo struct {
	Username   string
	Email      string
	SecretCode string
}

type UpdateVerifyEmailRepo struct {
	ID         int64
	SecretCode string
}

type VerifyEmailTxRepo struct {
	EmailId     int64
	SecreteCode string
}

type VerifyEmailTxRepoResult struct {
	User        *domain.User
	VerifyEmail *domain.VerifyEmail
}
