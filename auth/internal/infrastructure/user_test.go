package infrastructure

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) domain.User {
	t.Helper()

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := application.CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomUsername(),
		Email:          util.RandomEmail(),
	}

	user, err := repositories.User().CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return *user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestCreteUserSameUsername(t *testing.T) {
	user := createRandomUser(t)

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := application.CreateUserParams{
		Username:       user.Username,
		HashedPassword: hashedPassword,
		FullName:       util.RandomUsername(),
		Email:          util.RandomEmail(),
	}

	createdUser, err := repositories.User().CreateUser(context.Background(), arg)
	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrUsernameAlreadyExist)
	require.Nil(t, createdUser)
}

func TestCreteUserSameEmail(t *testing.T) {
	user := createRandomUser(t)

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := application.CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomUsername(),
		Email:          user.Email,
	}

	createdUser, err := repositories.User().CreateUser(context.Background(), arg)
	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrEmailAlreadyExist)
	require.Nil(t, createdUser)
}

func TestCreateUserTx(t *testing.T) {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := application.CreateUserTxParams{
		CreateUserParams: application.CreateUserParams{
			Username:       util.RandomUsername(),
			HashedPassword: hashedPassword,
			FullName:       util.RandomUsername(),
			Email:          util.RandomEmail(),
		},
		AfterCreate: func(user domain.User) error { return nil },
	}

	result, err := repositories.User().CreateUserTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result.User)

	user := result.User
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
}

func TestCreatUserTxRollBack(t *testing.T) {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := application.CreateUserTxParams{
		CreateUserParams: application.CreateUserParams{
			Username:       util.RandomUsername(),
			HashedPassword: hashedPassword,
			FullName:       util.RandomUsername(),
			Email:          util.RandomEmail(),
		},
		AfterCreate: func(user domain.User) error { return fmt.Errorf("error inside transaction") },
	}

	_, err = repositories.User().CreateUserTx(context.Background(), arg)
	require.Error(t, err)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := repositories.User().GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUserNotFound(t *testing.T) {
	user, err := repositories.User().GetUser(context.Background(), "not found")
	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrUserNotFound)
	require.Nil(t, user)
}

func TestUpdateUserOnlyFullname(t *testing.T) {
	oldUser := createRandomUser(t)
	newFullName := util.RandomUsername()

	updatedUser, err := repositories.User().UpdateUser(context.Background(), application.UpdateUserParams{
		Username: oldUser.Username,
		FullName: &newFullName,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, newFullName, updatedUser.FullName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)
	newEmail := util.RandomEmail()

	updatedUser, err := repositories.User().UpdateUser(context.Background(), application.UpdateUserParams{
		Username: oldUser.Username,
		Email:    &newEmail,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)
	newHashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	updatedUser, err := repositories.User().UpdateUser(context.Background(), application.UpdateUserParams{
		Username:       oldUser.Username,
		HashedPassword: &newHashedPassword,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)

	newFullName := util.RandomUsername()
	newEmail := util.RandomEmail()
	newHashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	updatedUser, err := repositories.User().UpdateUser(context.Background(), application.UpdateUserParams{
		Username:       oldUser.Username,
		FullName:       &newFullName,
		Email:          &newEmail,
		HashedPassword: &newHashedPassword,
	})
	require.NoError(t, err)

	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)

	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newFullName, updatedUser.FullName)
	require.Equal(t, newEmail, updatedUser.Email)
}
