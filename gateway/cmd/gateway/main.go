package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/pkg/token"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/gateway"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/middleware"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)

	startMetrics()
	startServer(ctx, waitGroup, config)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func startMetrics() {
	prometheus.MustRegister(middleware.HttpRequestsTotal)
	prometheus.MustRegister(middleware.HttpRequestDuration)
}

func startServer(ctx context.Context, waitGroup *errgroup.Group, config util.Config) {
	gw, settings, err := gateway.NewGateway(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create gateway")
	}

	jwtToken, err := token.NewJwtToken(config.TokenSecret)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot initiate jwt token")
	}

	middleware := middleware.NewMiddleware(config, settings, jwtToken)

	routes := chi.NewRouter()
	routes.Get("/metrics", promhttp.Handler().ServeHTTP)

	routes.With(
		middleware.RecoverPanic,
		middleware.Metric,
		middleware.AllowCors,
		middleware.Logger,
		middleware.RateLimit,
		middleware.Authenticate,
	).Route("/", func(r chi.Router) {
		r.Mount("/", gw)
	})

	srv := &http.Server{
		Addr:    config.ServerAddress,
		Handler: routes,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gateway server at %s", config.ServerAddress)
		err = srv.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			log.Fatal().Err(err).Msg("gateway server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown gateway server")
		srv.Shutdown(context.Background())
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("failed to shutdown gateway server")
			return err
		}

		log.Info().Msg("HTTP gateway server was stopped")
		return nil
	})
}
