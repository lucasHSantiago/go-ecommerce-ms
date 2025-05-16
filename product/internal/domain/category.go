package domain

import (
	"errors"
	"time"
)

var ErrCategoryNotFound = errors.New("category not found")
var ErrReadCategory = errors.New("failed to read category")
var ErrCreateCategory = errors.New("failed to create category")
var ErrUpdateCategory = errors.New("failed to update category")

type Category struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
