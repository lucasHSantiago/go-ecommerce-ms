package infra

import (
	"context"
	"testing"

	"github.com/lucasHSantiago/go-ecommerce-ms/product/internal/domain"
	"github.com/lucasHSantiago/go-ecommerce-ms/product/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) *domain.Category {
	t.Helper()

	arg := CreateCategoryParams{
		Name: util.RandomCategory(),
	}

	category, err := repositories.Category().CreateCategory(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, category)
	require.Equal(t, arg.Name, category.Name)
	require.NotEmpty(t, category.ID)
	require.NotZero(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategoryAll(t *testing.T) {
	createRandomCategory(t)

	arg := ListCategoryParams{
		Limit:  3,
		Offset: 0,
	}

	categories, err := repositories.Category().GetCategoryAll(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, categories)
	require.Equal(t, len(categories), 3)

	for _, c := range categories {
		require.NotEmpty(t, c.ID)
		require.NotEmpty(t, c.Name)
		require.NotEmpty(t, c.CreatedAt)
	}
}

func TestGetCategoryById(t *testing.T) {
	category := createRandomCategory(t)

	readCategory, err := repositories.Category().GetCategoryById(context.Background(), category.ID)
	require.NoError(t, err)

	require.Equal(t, category.ID, readCategory.ID)
	require.Equal(t, category.Name, readCategory.Name)
	require.Equal(t, category.CreatedAt, readCategory.CreatedAt)
}

func TestGetCategoryById_NotFound(t *testing.T) {
	_, err := repositories.Category().GetCategoryById(context.Background(), 9999999)
	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrCategoryNotFound)
}

func TestUpdateCategory(t *testing.T) {
	category := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:   category.ID,
		Name: util.RandomCategory(),
	}

	updatedCategory, err := repositories.Category().UpdateCategory(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, category.ID, updatedCategory.ID)
	require.Equal(t, category.CreatedAt, updatedCategory.CreatedAt)
	require.Equal(t, arg.Name, updatedCategory.Name)
}
