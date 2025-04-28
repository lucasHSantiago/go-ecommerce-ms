package infra

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/application/port"
	"github.com/lucasHSantiago/go-ecommerce-ms/auth/internal/util"
)

type testRepositories struct {
	connPool    *pgxpool.Pool
	user        port.UserRepository
	verifyEmail port.VerifyEmailRepository
}

func (r *testRepositories) User() port.UserRepository {
	if r.user == nil {
		r.user = NewUserRepository(r.connPool)
	}

	return r.user
}

func (r *testRepositories) VerifyEmail() port.VerifyEmailRepository {
	if r.verifyEmail == nil {
		r.verifyEmail = NewVerifyEmailRepository(r.connPool)
	}

	return r.verifyEmail
}

var repositories testRepositories

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	repositories = testRepositories{
		connPool: connPool,
	}

	os.Exit(m.Run())
}
