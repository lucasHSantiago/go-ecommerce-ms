package gapi

import (
	"fmt"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/token"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
)

type Server struct {
	gen.UnimplementedAuthServer
	config          util.Config
	userApplication UserApplication
	tokenMaker      token.Maker
	taskDistributor application.TaskDistributor
}

func NewServer(config util.Config, userApplication UserApplication, taskDistributor application.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		userApplication: userApplication,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
