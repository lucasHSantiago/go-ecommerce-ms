package infra

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucasHSantiago/go-ecommerce-ms/product/internal/util"
)

type testRepositories struct {
	connPool *pgxpool.Pool
	category *CategoryRepository
}

func (r *testRepositories) Category() *CategoryRepository {
	if r.category == nil {
		r.category = NewCategoryRepository(r.connPool)
	}

	return r.category
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
