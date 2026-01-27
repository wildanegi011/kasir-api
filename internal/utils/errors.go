package utils

import "errors"

var (
	ErrEmptyDatabaseURL = errors.New("DATABASE_URL is not set")

	ErrProductNotFound  = errors.New("product not found")
	ErrCategoryNotFound = errors.New("category not found")
)
