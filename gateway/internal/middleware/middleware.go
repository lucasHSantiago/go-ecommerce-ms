package middleware

import (
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/pkg/token"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/gateway"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/util"
)

type TokenVerifier interface {
	VerifyToken(token string) (*token.Payload, error)
}

type Middleware struct {
	config          util.Config
	gatewaySettings *gateway.GatewaySettings
	tokenVerifier   TokenVerifier
}

func NewMiddleware(config util.Config, gatewaySettings *gateway.GatewaySettings, tokenVerifier TokenVerifier) *Middleware {
	return &Middleware{config, gatewaySettings, tokenVerifier}
}
