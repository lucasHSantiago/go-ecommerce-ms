package infrastructure

import (
	"context"
	"fmt"
	"testing"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomVerifyEmail(t *testing.T) domain.VerifyEmail {
	t.Helper()

	user := createRandomUser(t)

	arg := port.CreateVerifyEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: util.RandomString(32),
	}

	verifyEmail, err := repositories.VerifyEmail().CreateVerifyEmail(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, verifyEmail)

	require.Equal(t, arg.Username, verifyEmail.Username)
	require.Equal(t, arg.Email, verifyEmail.Email)
	require.Equal(t, arg.SecretCode, verifyEmail.SecretCode)

	require.False(t, verifyEmail.IsUsed)
	require.NotZero(t, verifyEmail.CreatedAt)
	require.NotZero(t, verifyEmail.ExpiredAt)

	return *verifyEmail
}

func TestCreateVerifyEmail(t *testing.T) {
	createRandomVerifyEmail(t)
}

func TestCreateVerifyEmailUsernameInvalid(t *testing.T) {
	arg := port.CreateVerifyEmailParams{
		Username:   "username invalid",
		Email:      "email invalid",
		SecretCode: util.RandomString(32),
	}

	verifyEmail, err := repositories.VerifyEmail().CreateVerifyEmail(context.Background(), arg)

	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrUserNotFound)
	require.Nil(t, verifyEmail)
}

func TestUpdateVerifyEmail(t *testing.T) {
	verifyEmail := createRandomVerifyEmail(t)

	arg := port.UpdateVerifyEmailParams{
		ID:         verifyEmail.ID,
		SecretCode: verifyEmail.SecretCode,
	}

	fmt.Print(arg.ID)

	updatedVerifyEmail, err := repositories.VerifyEmail().UpdateVerifyEmail(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, updatedVerifyEmail.SecretCode, verifyEmail.SecretCode)
	require.Equal(t, updatedVerifyEmail.SecretCode, arg.SecretCode)
	require.True(t, updatedVerifyEmail.IsUsed)
}
