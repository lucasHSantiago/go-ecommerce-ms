package params

type CreateVerifyEmailRepo struct {
	Username   string
	Email      string
	SecretCode string
}

type UpdateVerifyEmailRepo struct {
	ID         int64
	SecretCode string
}
