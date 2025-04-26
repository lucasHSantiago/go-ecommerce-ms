package infrastructure

// func createRandomVerifyEmail(t *testing.T) VerifyEmail {
// 	t.Helper()
//
// 	user := createRandomUser(t)
//
// 	arg := CreateVerifyEmailParams{
// 		Username:   user.Username,
// 		Email:      user.Email,
// 		SecretCode: util.RandomString(32),
// 	}
//
// 	verifyEmail, err := testQueries.CreateVerifyEmail(context.Background(), arg)
//
// 	require.NoError(t, err)
// 	require.NotEmpty(t, verifyEmail)
//
// 	require.Equal(t, arg.Username, verifyEmail.Username)
// 	require.Equal(t, arg.Email, verifyEmail.Email)
// 	require.Equal(t, arg.SecretCode, verifyEmail.SecretCode)
//
// 	require.False(t, verifyEmail.IsUsed)
// 	require.NotZero(t, verifyEmail.CreatedAt)
// 	require.NotZero(t, verifyEmail.ExpiredAt)
//
// 	return verifyEmail
// }
//
// func TestCreateVerifyEmail(t *testing.T) {
// 	createRandomVerifyEmail(t)
// }
//
// func TestUpdateVerifyEmail(t *testing.T) {
// 	verifyEmail := createRandomVerifyEmail(t)
//
// 	arg := UpdateVerifyEmailParams{
// 		ID:         verifyEmail.ID,
// 		SecretCode: verifyEmail.SecretCode,
// 	}
//
// 	fmt.Print(arg.ID)
//
// 	updatedVerifyEmail, err := testQueries.UpdateVerifyEmail(context.Background(), arg)
// 	require.NoError(t, err)
//
// 	require.Equal(t, updatedVerifyEmail.SecretCode, verifyEmail.SecretCode)
// 	require.Equal(t, updatedVerifyEmail.SecretCode, arg.SecretCode)
// 	require.True(t, updatedVerifyEmail.IsUsed)
// }
