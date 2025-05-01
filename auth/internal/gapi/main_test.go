package gapi

import (
	"testing"
	"time"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, userApplication UserApplication) *Server {
	config := util.Config{
		TokenSecretKey:      util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, userApplication, nil)
	require.NoError(t, err)

	return server
}
