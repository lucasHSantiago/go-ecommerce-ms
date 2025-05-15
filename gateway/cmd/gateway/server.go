package main

import (
	"context"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/pkg/token"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/api"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/gateway"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/middleware"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	config util.Config
}

func NewServer() *Server {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return &Server{config}
}

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func (s *Server) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)

	s.startMetrics()
	s.startApi(ctx, waitGroup)

	err := waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func (s *Server) startMetrics() {
	prometheus.MustRegister(middleware.HttpRequestsTotal)
	prometheus.MustRegister(middleware.HttpRequestDuration)
}

func (s *Server) startApi(ctx context.Context, waitGroup *errgroup.Group) {
	gw, settings, err := gateway.NewGateway(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create gateway")
	}

	routes := s.setupRoutes(gw, settings)

	srv := &http.Server{
		Addr:    s.config.ServerAddress,
		Handler: routes,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gateway server at %s", s.config.ServerAddress)
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

func (s *Server) setupRoutes(gw http.Handler, settings *gateway.GatewaySettings) *chi.Mux {
	jwtToken, err := token.NewJwtToken(s.config.TokenSecret)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot initiate jwt token")
	}

	middleware := middleware.NewMiddleware(s.config, settings, jwtToken)

	routes := chi.NewRouter()

	routes.Use(middleware.AllowCors)

	fsys, err := fs.Sub(api.StaticSwaggerFS, "swagger")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot initiate swagger file server")
	}

	routes.Handle("/swagger/*", http.StripPrefix("/swagger", http.FileServer(http.FS(fsys))))

	routes.Get("/metrics", promhttp.Handler().ServeHTTP)

	routes.With(
		middleware.RecoverPanic,
		middleware.Metric,
		middleware.Logger,
		middleware.RateLimit,
		middleware.Authenticate,
	).Route("/", func(r chi.Router) {
		r.Mount("/", gw)
	})

	return routes
}
