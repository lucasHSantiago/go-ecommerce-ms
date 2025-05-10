package middleware

import "github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/util"

type Middleware struct {
	config util.Config
}

func NewMiddleware(config util.Config) *Middleware {
	return &Middleware{config}
}
