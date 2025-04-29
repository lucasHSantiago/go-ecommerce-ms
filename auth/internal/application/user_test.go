package application

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port"
	mockdb "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port/mock"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/gapi/service"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/validator"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/worker"
	mockwk "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/worker/mock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsTxMatcher struct {
	arg      port.CreateUserTxParams
	password string
	user     domain.User
}

func (expected eqCreateUserParamsTxMatcher) Matches(x any) bool {
	actualArg, ok := x.(port.CreateUserTxParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(expected.password, actualArg.HashedPassword)
	if err != nil {
		return false
	}

	expected.arg.HashedPassword = actualArg.HashedPassword
	if !reflect.DeepEqual(expected.arg.CreateUserParams, actualArg.CreateUserParams) {
		return false
	}

	err = actualArg.AfterCreate(expected.user)

	return err == nil
}

func (e eqCreateUserParamsTxMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserTxParams(arg port.CreateUserTxParams, password string, user domain.User) gomock.Matcher {
	return eqCreateUserParamsTxMatcher{arg, password, user}
}

func randomUser(t *testing.T) (domain.User, string) {
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

	return user, password
}

func TestCreateUserUseCase(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		arg           service.CreateUserParams
		buildStubs    func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor)
		checkResponse func(t *testing.T, res *domain.User, err error)
	}{
		{
			name: "OK",
			arg: service.CreateUserParams{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				arg := port.CreateUserTxParams{
					CreateUserParams: port.CreateUserParams{
						Username: user.Username,
						FullName: user.FullName,
						Email:    user.Email,
					},
				}

				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), EqCreateUserTxParams(arg, password, user)).
					Times(1).
					Return(port.CreateUserTxResult{User: user}, nil)

				taskPayload := &worker.PayloadSendVerifyEmail{
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
			arg: service.CreateUserParams{
				Username: "",
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidUsername",
			arg: service.CreateUserParams{
				Username: "invalid 123",
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredFullName",
			arg: service.CreateUserParams{
				Username: user.Username,
				FullName: "",
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidFullName",
			arg: service.CreateUserParams{
				Username: user.Username,
				FullName: "Invalid123",
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredEmail",
			arg: service.CreateUserParams{
				Username: user.Username,
				FullName: user.FullName,
				Email:    "",
				Password: password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidEmail",
			arg: service.CreateUserParams{
				Username: user.Username,
				FullName: user.FullName,
				Email:    "invalid",
				Password: password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredPassword",
			arg: service.CreateUserParams{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: "",
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)

				taskDistributor.EXPECT().
					DistributeTaskSendVerifyEmail(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "InternalError",
			arg: service.CreateUserParams{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository, taskDistributor *mockwk.MockTaskDistributor) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(port.CreateUserTxResult{}, sql.ErrConnDone)

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
			defer repositoryCtrl.Finish()
			userRespository := mockdb.NewMockUserRepository(repositoryCtrl)

			distributorCtrl := gomock.NewController(t)
			defer distributorCtrl.Finish()
			taskDistrubutor := mockwk.NewMockTaskDistributor(distributorCtrl)

			tc.buildStubs(userRespository, taskDistrubutor)

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
		arg           service.UpdateUserParams
		buildStubs    func(userRepository *mockdb.MockUserRepository)
		checkResponse func(t *testing.T, updatedUser *domain.User, err error)
	}{
		{
			name: "OK",
			arg: service.UpdateUserParams{
				Username: user.Username,
				FullName: &newFullname,
				Email:    &newEmail,
			},
			buildStubs: func(store *mockdb.MockUserRepository) {
				arg := port.UpdateUserParams{
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
			arg: service.UpdateUserParams{
				Username: "",
				FullName: &user.FullName,
				Email:    &user.Email,
				Password: &password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidUsername",
			arg: service.UpdateUserParams{
				Username: "invalid 123",
				FullName: &user.FullName,
				Email:    &user.Email,
				Password: &password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredFullName",
			arg: service.UpdateUserParams{
				Username: user.Username,
				FullName: new(string),
				Email:    &user.Email,
				Password: &password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidFullName",
			arg: service.UpdateUserParams{
				Username: user.Username,
				FullName: &invalidFullname,
				Email:    &user.Email,
				Password: &password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredEmail",
			arg: service.UpdateUserParams{
				Username: user.Username,
				FullName: &user.FullName,
				Email:    new(string),
				Password: &password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidEmail",
			arg: service.UpdateUserParams{
				Username: user.Username,
				FullName: &user.FullName,
				Email:    &invalidEmail,
				Password: &password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "RequiredPassword",
			arg: service.UpdateUserParams{
				Username: user.Username,
				FullName: &user.FullName,
				Email:    &user.Email,
				Password: new(string),
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *domain.User, err error) {
				require.Error(t, err)
				_, ok := err.(*validator.ValidationErrors)
				require.True(t, ok)
			},
		},
		{
			name: "InternalError",
			arg: service.UpdateUserParams{
				Username: user.Username,
				FullName: &user.FullName,
				Email:    &user.Email,
				Password: &password,
			},
			buildStubs: func(userRepository *mockdb.MockUserRepository) {
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
			defer repositoryCtrl.Finish()
			userRespository := mockdb.NewMockUserRepository(repositoryCtrl)

			tc.buildStubs(userRespository)

			userApplication := NewUserApplication(userRespository, nil, nil, nil, nil)
			res, err := userApplication.Update(context.Background(), tc.arg)
			tc.checkResponse(t, res, err)
		})
	}
}
