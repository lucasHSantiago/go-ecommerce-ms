package infra

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/lucasHSantiago/go-ecommerce-ms/product/internal/domain"
	"github.com/rs/zerolog/log"
)

type CategoryRepository struct {
	connPool DBTX
}

func NewCategoryRepository(connPool DBTX) *CategoryRepository {
	return &CategoryRepository{connPool}
}

const selectCategoryById = `
	SELECT id, name, created_at FROM categories WHERE id = $1;
	`

func (c *CategoryRepository) GetCategoryById(ctx context.Context, id int64) (*domain.Category, error) {
	rows, _ := c.connPool.Query(ctx, selectCategoryById, id)

	category, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Category])
	if err != nil {
		return nil, getCategoryError(err, domain.ErrReadCategory)
	}

	return category, nil
}

const selectCategoryAll = `
	SELECT id, name, created_at
	FROM categories
	ORDER by created_at desc
	LIMIT $1
	OFFSET $2;
	`

type ListCategoryParams struct {
	Limit  int32
	Offset int32
}

func (c *CategoryRepository) GetCategoryAll(ctx context.Context, arg ListCategoryParams) ([]*domain.Category, error) {
	rows, _ := c.connPool.Query(ctx, selectCategoryAll, arg.Limit, arg.Offset)

	category, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Category])
	if err != nil {
		return nil, getCategoryError(err, domain.ErrReadCategory)
	}

	return category, nil
}

type CreateCategoryParams struct {
	Name string
}

const createCategory = `
	INSERT INTO categories (name, created_at) VALUES ($1, now()) RETURNING id, name, created_at;
	`

func (c *CategoryRepository) CreateCategory(ctx context.Context, arg CreateCategoryParams) (*domain.Category, error) {
	rows, _ := c.connPool.Query(ctx, createCategory, arg.Name)

	category, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Category])
	if err != nil {
		return nil, getCategoryError(err, domain.ErrCreateCategory)
	}

	return category, nil
}

type UpdateCategoryParams struct {
	ID   int64
	Name string
}

const updateCategory = `
	UPDATE categories SET name = $1 WHERE id = $2 RETURNING id, name, created_at;
	`

func (c *CategoryRepository) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (*domain.Category, error) {
	args := []any{
		arg.Name,
		arg.ID,
	}

	rows, _ := c.connPool.Query(ctx, updateCategory, args...)

	category, err := pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[domain.Category])
	if err != nil {
		return nil, getCategoryError(err, domain.ErrUpdateCategory)
	}

	return category, nil
}

func getCategoryError(err error, defaultReturn error) error {
	if errors.Is(err, ErrRecordNotFound) {
		return domain.ErrCategoryNotFound
	}

	log.Error().Err(err)
	return defaultReturn
}
