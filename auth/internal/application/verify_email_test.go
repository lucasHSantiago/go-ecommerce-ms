package application

import (
	"context"
	"testing"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/golang/mock/gomock"
	mockdb "github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/mock"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/infra"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/stretchr/testify/require"
)

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

func TestVerifyEmailUseCase(t *testing.T) {
	user, _ := randomUser(t)
	verifyEmail := randomVerifyEmail(user)

	testCases := []struct {
		name          string
		arg           VerifyEmail
		buildMocks    func(arg VerifyEmail, verifyEmailRepository *mockdb.MockVerifyEmailRepository)
		checkResponse func(t *testing.T, res *VerifyEmailResult, err error)
	}{
		{
			name: "OK",
			arg: VerifyEmail{
				EmailId:    verifyEmail.ID,
				SecretCode: verifyEmail.SecretCode,
			},
			buildMocks: func(arg VerifyEmail, verifyEmailRepository *mockdb.MockVerifyEmailRepository) {
				txArg := infra.VerifyEmailTx{
					EmailId:     arg.EmailId,
					SecreteCode: arg.SecretCode,
				}

				txRes := infra.VerifyEmailTxResult{
					User:        user,
					VerifyEmail: verifyEmail,
				}

				verifyEmailRepository.EXPECT().
					VerifyEmailTx(gomock.Any(), txArg).
					Times(1).
					Return(txRes, nil)
			},
			checkResponse: func(t *testing.T, res *VerifyEmailResult, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, res.User.Username, user.Username)
				require.Equal(t, res.VerifyEmail.ID, verifyEmail.ID)
			},
		},
		{
			name: "RequiredId",
			arg: VerifyEmail{
				EmailId:    0,
				SecretCode: verifyEmail.SecretCode,
			},
			buildMocks: func(arg VerifyEmail, verifyEmailRepository *mockdb.MockVerifyEmailRepository) {
				verifyEmailRepository.EXPECT().
					VerifyEmailTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *VerifyEmailResult, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
		{
			name: "InvalidSecretCode",
			arg: VerifyEmail{
				EmailId:    verifyEmail.ID,
				SecretCode: "invalid",
			},
			buildMocks: func(arg VerifyEmail, verifyEmailRepository *mockdb.MockVerifyEmailRepository) {
				verifyEmailRepository.EXPECT().
					VerifyEmailTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *VerifyEmailResult, err error) {
				require.Error(t, err)
				require.Nil(t, res)
				_, ok := err.(validation.Errors)
				require.True(t, ok)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repositoryCtrl := gomock.NewController(t)
			verifyEmailRespository := mockdb.NewMockVerifyEmailRepository(repositoryCtrl)
			tc.buildMocks(tc.arg, verifyEmailRespository)

			verifyEmailApplication := NewVerifyEmailApplication(verifyEmailRespository)
			res, err := verifyEmailApplication.VerifyEmail(context.Background(), tc.arg)
			tc.checkResponse(t, res, err)
		})
	}
}
