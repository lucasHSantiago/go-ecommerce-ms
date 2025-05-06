package main

import (
	"context"
	"net/http"

	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/gateway"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/util"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	//@TODO: format logs on debug

	gw, err := gateway.NewGateway(context.Background(), config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create gateway")
	}

	log.Info().Msg("gateway created")

	mux := http.NewServeMux()
	mux.Handle("/", gw)

	log.Info().Msg("mux created")

	log.Info().Msg("Server started on localhost:9090")

	err = http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start http server")
	}
}
