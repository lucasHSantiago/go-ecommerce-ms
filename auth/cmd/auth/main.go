package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/gapi"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/infra"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/mail"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/worker"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/pkg/token"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/proto/gen"
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

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	verifyEmailRepository := infra.NewVerifyEmailRepository(connPool)
	userRepository := infra.NewUserRepository(connPool)

	waitGroup, ctx := errgroup.WithContext(ctx)
	runTaskProcessor(ctx, waitGroup, userRepository, verifyEmailRepository, config)
	runGrpcServer(ctx, connPool, waitGroup, userRepository, verifyEmailRepository, config)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runGrpcServer(ctx context.Context, connPool *pgxpool.Pool, waitGroup *errgroup.Group, userRepository application.UserRepository, verifyEmailRepository application.VerifyEmailRepository, config util.Config) {
	userApplication := newUserApplication(connPool, userRepository, &config)
	verifyEmailApplication := newVerifyEmailApplication(verifyEmailRepository)
	server := gapi.NewAuthServer(userApplication, verifyEmailApplication)

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	gen.RegisterAuthServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gRPC server at %s\n", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}

			log.Error().Err(err).Msg("gRPC failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown gRPC server")
		grpcServer.GracefulStop()
		log.Info().Msg("gRPC server was stopped")

		return nil
	})
}

func runTaskProcessor(ctx context.Context, waitGroup *errgroup.Group, userRepository application.UserRepository, verifyEmailRepository application.VerifyEmailRepository, config util.Config) {
	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	mailer := mail.NewMailTrappSender(config.EmailSenderUsername, config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, userRepository, verifyEmailRepository, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown task processor")
		taskProcessor.Shutdown()
		log.Info().Msg("task processor was stopped")
		return nil
	})
}

func newUserApplication(connPool *pgxpool.Pool, userRepository application.UserRepository, config *util.Config) gapi.UserApplication {
	tokenMaker, err := token.NewJwtToken(config.TokenSecretKey)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create token maker")
	}

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	sessionRepository := infra.NewSessionRepository(connPool)

	return application.NewUserApplication(userRepository, sessionRepository, taskDistributor, tokenMaker, config)
}

func newVerifyEmailApplication(verifyEmailRepository application.VerifyEmailRepository) gapi.VerifyEmailApplication {
	return application.NewVerifyEmailApplication(verifyEmailRepository)
}
