package application

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/distributor"
	mockwk "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/distributor/mock"
	mockdb "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/mock"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/token"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserUseCase(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		arg           params.CreateUserApp
		buildMocks    func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor)
		checkResponse func(t *testing.T, res *domain.User, err error)
	}{
		{
			name: "OK",
			arg: params.CreateUserApp{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				arg := params.CreateUserTxRepo{
					CreateUserRepo: params.CreateUserRepo{
						Username: user.Username,
						FullName: user.FullName,
						Email:    user.Email,
					},
				}

				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), EqCreateUserTxParams(arg, password, *user)).
					Times(1).
					Return(params.CreateUserTxRepoResult{User: *user}, nil)

				taskPayload := &distributor.PayloadSendVerifyEmail{
					Username: user.Username,
				}

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), taskPayload, gomock.Any()).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, user.Username, res.Username)
				require.Equal(t, user.FullName, res.FullName)
				require.Equal(t, user.Email, res.Email)
			},
		},
		{
			name: "RequiredUserName",
			arg: params.CreateUserApp{
				Username: "",
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidUsername",
			arg: params.CreateUserApp{
				Username: "invalid 123",
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredFullName",
			arg: params.CreateUserApp{
				Username: user.Username,
				FullName: "",
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidFullName",
			arg: params.CreateUserApp{
				Username: user.Username,
				FullName: "Invalid123",
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredEmail",
			arg: params.CreateUserApp{
				Username: user.Username,
				FullName: user.FullName,
				Email:    "",
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidEmail",
			arg: params.CreateUserApp{
				Username: user.Username,
				FullName: user.FullName,
				Email:    "invalid",
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredPassword",
			arg: params.CreateUserApp{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: "",
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InternalError",
			arg: params.CreateUserApp{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(params.CreateUserTxRepoResult{}, sql.ErrConnDone)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrConnDone)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repositoryCtrl := gomock.NewController(t)
			userRespository := mockdb.NewMockUserRepository(repositoryCtrl)

			distributorCtrl := gomock.NewController(t)
			taskDistrubutor := mockwk.NewMockTaskDistributor(distributorCtrl)

			tc.buildMocks(userRespository, taskDistrubutor)

			userApplication := NewUserApplication(userRespository, nil, taskDistrubutor, nil, nil)
			res, err := userApplication.Create(context.Background(), tc.arg)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestUpdateUserUseCase(t *testing.T) {
	user, password := randomUser(t)

	newFullname := util.RandomUsername()
	newEmail := util.RandomEmail()

	invalidEmail := "invalid-email"
	invalidFullname := "invalid123"

	tests := []struct {
		name          string
		arg           params.UpdateUserApp
		buildMocks    func(userRepository *mockdb.MockUserRepository)
		checkResponse func(t *testing.T, updatedUser *domain.User, err error)
	}{
		{
			name: "OK",
			arg: params.UpdateUserApp{
				Username: user.Username,
				FullName: &newFullname,
				Email:    &newEmail,
			},
			buildMocks: func(store *mockdb.MockUserRepository) {
				arg := params.UpdateUserRepo{
					Username: user.Username,
					FullName: &newFullname,
					Email:    &newEmail,
				}

				updatedUser := &domain.User{
					Username:          user.Username,
					HashedPassword:    user.HashedPassword,
					FullName:          newFullname,
					Email:             newEmail,
					PasswordChangedAt: user.PasswordChangedAt,
					CreatedAt:         user.CreatedAt,
					IsEmailVerified:   user.IsEmailVerified,
				}

				store.EXPECT().
					UpdateUser(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(updatedUser, nil)
			},
			checkResponse: func(t *testing.T, updatedUser *domain.User, err error) {
				require.NoError(t, err)
				require.NotNil(t, updatedUser)
				require.Equal(t, user.Username, updatedUser.Username)
				require.Equal(t, newFullname, updatedUser.FullName)
				require.Equal(t, newEmail, updatedUser.Email)
			},
		},
		{
			name: "RequiredUserName",
			arg: params.UpdateUserApp{
				Username: "",
				FullName: &user.FullName,
				Email:    &user.Email,
				Password: &password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidUsername",
			arg: params.UpdateUserApp{
				Username: "invalid 123",
				FullName: &user.FullName,
				Email:    &user.Email,
				Password: &password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredFullName",
			arg: params.UpdateUserApp{
				Username: user.Username,
				FullName: new(string),
				Email:    &user.Email,
				Password: &password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidFullName",
			arg: params.UpdateUserApp{
				Username: user.Username,
				FullName: &invalidFullname,
				Email:    &user.Email,
				Password: &password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredEmail",
			arg: params.UpdateUserApp{
				Username: user.Username,
				FullName: &user.FullName,
				Email:    new(string),
				Password: &password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidEmail",
			arg: params.UpdateUserApp{
				Username: user.Username,
				FullName: &user.FullName,
				Email:    &invalidEmail,
				Password: &password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredPassword",
			arg: params.UpdateUserApp{
				Username: user.Username,
				FullName: &user.FullName,
				Email:    &user.Email,
				Password: new(string),
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InternalError",
			arg: params.UpdateUserApp{
				Username: user.Username,
				FullName: &user.FullName,
				Email:    &user.Email,
				Password: &password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(&domain.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrConnDone)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repositoryCtrl := gomock.NewController(t)
			userRespository := mockdb.NewMockUserRepository(repositoryCtrl)

			tc.buildMocks(userRespository)

			userApplication := NewUserApplication(userRespository, nil, nil, nil, nil)
			res, err := userApplication.Update(context.Background(), tc.arg)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestLoginUserUseCase(t *testing.T) {
	user, password := randomUser(t)
	session := randomSession(t, user.Username)

	testCases := []struct {
		name          string
		arg           params.LoginUserApp
		buildMocks    func(userRepository *mockdb.MockUserRepository, sessionRepository *mockdb.MockSessionRepository)
		checkResponse func(t *testing.T, result *params.LoginUserAppResult, err error)
	}{
		{
			name: "OK",
			arg: params.LoginUserApp{
				Username: user.Username,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, sessionRepository *mockdb.MockSessionRepository) {
				userRepository.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)

				sessionRepository.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(t *testing.T, result *params.LoginUserAppResult, err error) {
				require.Equal(t, result.User, user)
				require.Equal(t, result.SessionId, session.ID)
				require.NotEmpty(t, result.AccessToken)
				require.NotEmpty(t, result.RefreshToken)
				require.True(t, result.AccessTokenExpiresAt.After(time.Now()))
				require.True(t, result.RefreshTokenExpiresAt.After(time.Now()))
			},
		},
		{
			name: "IncorrectPassword",
			arg: params.LoginUserApp{
				Username: user.Username,
				Password: "incorrect",
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, sessionRepository *mockdb.MockSessionRepository) {
				userRepository.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)

				sessionRepository.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, result *params.LoginUserAppResult, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, ErrInvalidLoginPassword)
				require.Nil(t, result)
			},
		},
		{
			name: "InvalidUsername",
			arg: params.LoginUserApp{
				Username: "invalid-user#1",
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, sessionRepository *mockdb.MockSessionRepository) {
				userRepository.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)

				sessionRepository.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, result *params.LoginUserAppResult, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "PasswordTooShort",
			arg: params.LoginUserApp{
				Username: user.Username,
				Password: "short",
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, sessionRepository *mockdb.MockSessionRepository) {
				userRepository.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)

				sessionRepository.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, result *params.LoginUserAppResult, err error) {
				require.Error(t, err)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userCtrl := gomock.NewController(t)
			userRepository := mockdb.NewMockUserRepository(userCtrl)

			sessionCtrl := gomock.NewController(t)
			sessionRepository := mockdb.NewMockSessionRepository(sessionCtrl)

			tc.buildMocks(userRepository, sessionRepository)

			tokenMaker, err := token.NewJWTMaker(util.RandomString(32))
			require.NoError(t, err)

			config := util.Config{
				AccessTokenDuration:  time.Minute,
				RefreshTokenDuration: time.Minute,
			}

			userApplication := NewUserApplication(userRepository, sessionRepository, nil, tokenMaker, &config)

			result, err := userApplication.Login(context.Background(), tc.arg)
			tc.checkResponse(t, result, err)
		})
	}
}

type eqCreateUserParamsTxMatcher struct {
	arg      params.CreateUserTxRepo
	password string
	user     domain.User
}

func (expected eqCreateUserParamsTxMatcher) Matches(x any) bool {
	actualArg, ok := x.(params.CreateUserTxRepo)
	if !ok {
		return false
	}

	err := util.CheckPassword(expected.password, actualArg.HashedPassword)
	if err != nil {
		return false
	}

	expected.arg.HashedPassword = actualArg.HashedPassword
	if !reflect.DeepEqual(expected.arg.CreateUserRepo, actualArg.CreateUserRepo) {
		return false
	}

	err = actualArg.AfterCreate(expected.user)

	return err == nil
}

func (e eqCreateUserParamsTxMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserTxParams(arg params.CreateUserTxRepo, password string, user domain.User) gomock.Matcher {
	return eqCreateUserParamsTxMatcher{arg, password, user}
}

func randomUser(t *testing.T) (*domain.User, string) {
	t.Helper()

	password := util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user := domain.User{
		Username:       util.RandomUsername(),
		Role:           domain.UserRole,
		HashedPassword: hashedPassword,
		FullName:       util.RandomUsername(),
		Email:          util.RandomEmail(),
	}

	return &user, password
}

func randomSession(t *testing.T, username string) *domain.Session {
	t.Helper()

	session := domain.Session{
		ID:           uuid.New(),
		Username:     username,
		RefreshToken: "",
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(time.Minute),
		CreatedAt:    time.Now().Add(-time.Minute),
	}

	return &session
}
