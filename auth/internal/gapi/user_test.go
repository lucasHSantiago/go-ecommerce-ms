package gapi

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application"
	mockdb "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/mock"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/params"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
				res := params.CreateUserTxRepoResult{
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
					Return(params.CreateUserTxRepoResult{}, domain.ErrCreateUser)
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
			server := newTestServer(t, userApplication)

			res, err := server.CreateUser(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
