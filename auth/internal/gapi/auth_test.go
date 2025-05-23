package gapi

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application"
	mockdb "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/mock"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/infra"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/pkg/token"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		req           *gen.CreateUserRequest
		buildMocks    func(userRepository *mockdb.MockUserRepository)
		checkResponse func(t *testing.T, res *gen.CreateUserResponse, err error)
	}{
		{
			name: "OK",
			req: &gen.CreateUserRequest{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				res := infra.CreateUserTxResult{
					User: domain.User{
						Username: user.Username,
						FullName: user.FullName,
						Email:    user.Email,
					},
				}

				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(res, nil)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotNil(t, res.User)

				userRes := res.GetUser()
				require.Equal(t, user.Username, userRes.Username)
				require.Equal(t, user.FullName, userRes.FullName)
				require.Equal(t, user.Email, userRes.Email)
			},
		},
		{
			name: "RequiredUserName",
			req: &gen.CreateUserRequest{
				Username: "",
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "InvalidUsername",
			req: &gen.CreateUserRequest{
				Username: "invalid 123",
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "RequiredFullName",
			req: &gen.CreateUserRequest{
				Username: user.Username,
				FullName: "",
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "InvalidFullName",
			req: &gen.CreateUserRequest{
				Username: user.Username,
				FullName: "invalid123",
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "RequiredEmail",
			req: &gen.CreateUserRequest{
				Username: user.Username,
				FullName: user.FullName,
				Email:    "",
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "InvalidEmail",
			req: &gen.CreateUserRequest{
				Username: user.Username,
				FullName: user.FullName,
				Email:    "invalid@",
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "RequiredPassword",
			req: &gen.CreateUserRequest{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: "",
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &gen.CreateUserRequest{
				Username: user.Username,
				FullName: user.FullName,
				Email:    user.Email,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					CreateUserTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(infra.CreateUserTxResult{}, domain.ErrCreateUser)
			},
			checkResponse: func(t *testing.T, res *gen.CreateUserResponse, err error) {
				require.Error(t, err)
				require.Nil(t, res)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repositoryCtrl := gomock.NewController(t)
			userRespository := mockdb.NewMockUserRepository(repositoryCtrl)

			tc.buildMocks(userRespository)

			userApplication := application.NewUserApplication(userRespository, nil, nil, nil, nil)
			server := NewAuthServer(userApplication, nil)

			res, err := server.CreateUser(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestUpdateUserAPI(t *testing.T) {
	user, _ := randomUser(t)

	newName := util.RandomUsername()
	newEmail := util.RandomEmail()
	invalidEmail := "invalid-email"

	testCases := []struct {
		name          string
		req           *gen.UpdateUserRequest
		buildMocks    func(userRepository *mockdb.MockUserRepository)
		checkResponse func(t *testing.T, res *gen.UpdateUserResponse, err error)
	}{
		{
			name: "OK",
			req: &gen.UpdateUserRequest{
				Username: user.Username,
				FullName: &newName,
				Email:    &newEmail,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				updatedUser := &domain.User{
					Username:          user.Username,
					HashedPassword:    user.HashedPassword,
					FullName:          newName,
					Email:             newEmail,
					PasswordChangedAt: user.PasswordChangedAt,
					CreatedAt:         user.CreatedAt,
					IsEmailVerified:   user.IsEmailVerified,
				}

				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(updatedUser, nil)
			},
			checkResponse: func(t *testing.T, res *gen.UpdateUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				updatedUser := res.GetUser()
				require.Equal(t, user.Username, updatedUser.Username)
				require.Equal(t, newName, updatedUser.FullName)
				require.Equal(t, newEmail, updatedUser.Email)
			},
		},
		{
			name: "UserNotFound",
			req: &gen.UpdateUserRequest{
				Username: user.Username,
				FullName: &newName,
				Email:    &newEmail,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(&domain.User{}, domain.ErrUserNotFound)
			},
			checkResponse: func(t *testing.T, res *gen.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.NotFound, st.Code())
			},
		},
		{
			name: "InvalidEmail",
			req: &gen.UpdateUserRequest{
				Username: user.Username,
				FullName: &newName,
				Email:    &invalidEmail,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository) {
				userRepository.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *gen.UpdateUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repositoryCtrl := gomock.NewController(t)
			userRespository := mockdb.NewMockUserRepository(repositoryCtrl)

			tc.buildMocks(userRespository)

			userApplication := application.NewUserApplication(userRespository, nil, nil, nil, nil)
			server := NewAuthServer(userApplication, nil)

			res, err := server.UpdateUser(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestLoginUserAPI(t *testing.T) {
	user, password := randomUser(t)
	session := randomSession(t, user.Username)

	testCases := []struct {
		name          string
		req           *gen.LoginUserRequest
		buildMocks    func(userRepository *mockdb.MockUserRepository, sessionRepository *mockdb.MockSessionRepository)
		checkResponse func(t *testing.T, res *gen.LoginUserResponse, err error)
	}{
		{
			name: "OK",
			req: &gen.LoginUserRequest{
				Username: user.Username,
				Password: password,
			},
			buildMocks: func(userRepository *mockdb.MockUserRepository, sessionRepository *mockdb.MockSessionRepository) {
				userRepository.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				sessionRepository.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(t *testing.T, res *gen.LoginUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				loginUser := res.GetUser()
				require.Equal(t, user.Username, loginUser.Username)
				require.NotNil(t, res.AccessToken)
				require.NotNil(t, res.RefreshToken)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repositoryCtrl := gomock.NewController(t)
			userRespository := mockdb.NewMockUserRepository(repositoryCtrl)

			sessionCtrl := gomock.NewController(t)
			sessionRepository := mockdb.NewMockSessionRepository(sessionCtrl)

			tc.buildMocks(userRespository, sessionRepository)

			config := util.Config{
				AccessTokenDuration:  time.Minute,
				RefreshTokenDuration: time.Minute,
			}

			tokenMaker, err := token.NewJwtToken(util.RandomString(32))
			require.NoError(t, err)

			userApplication := application.NewUserApplication(userRespository, sessionRepository, nil, tokenMaker, &config)
			server := NewAuthServer(userApplication, nil)

			res, err := server.LoginUser(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}

func TestVerifyEmailAPI(t *testing.T) {
	user, _ := randomUser(t)
	verifyEmail := randomVerifyEmail(user)

	testCases := []struct {
		name          string
		req           *gen.VerifyEmailRequest
		buildMocks    func(verifyEmailRepository *mockdb.MockVerifyEmailRepository)
		checkResponse func(t *testing.T, res *gen.VerifyEmailResponse, err error)
	}{
		{
			name: "OK",
			req: &gen.VerifyEmailRequest{
				EmailId:    verifyEmail.ID,
				SecretCode: verifyEmail.SecretCode,
			},
			buildMocks: func(verifyEmailRepository *mockdb.MockVerifyEmailRepository) {
				user.IsEmailVerified = true

				txRes := infra.VerifyEmailTxResult{
					User:        user,
					VerifyEmail: verifyEmail,
				}

				verifyEmailRepository.EXPECT().
					VerifyEmailTx(gomock.Any(), gomock.Any()).
					Times(1).
					Return(txRes, nil)
			},
			checkResponse: func(t *testing.T, res *gen.VerifyEmailResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.True(t, res.IsVerified)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repositoryCtrl := gomock.NewController(t)
			verifyEmailRepository := mockdb.NewMockVerifyEmailRepository(repositoryCtrl)

			tc.buildMocks(verifyEmailRepository)

			verifyEmailApplication := application.NewVerifyEmailApplication(verifyEmailRepository)
			server := NewAuthServer(nil, verifyEmailApplication)

			res, err := server.VerifyEmail(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
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

func randomVerifyEmail(user *domain.User) *domain.VerifyEmail {
	return &domain.VerifyEmail{
		ID:         1,
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: util.RandomString(128),
		IsUsed:     false,
		CreatedAt:  time.Now(),
		ExpiredAt:  time.Now().Add(time.Minute),
	}
}
