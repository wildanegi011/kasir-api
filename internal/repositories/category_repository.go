package repositories

import (
	"context"
	"database/sql"
	domain "kasir-api/internal/domains"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context, page int, pageSize int) ([]domain.Category, int, error)
	GetCategoryByID(ctx context.Context, id int) (*domain.Category, error)
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	UpdateCategory(ctx context.Context, id int, category *domain.Category) (*domain.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

type CategoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (p *CategoryRepositoryImpl) GetCategories(ctx context.Context, page int, pageSize int) ([]domain.Category, int, error) {
	var total int
	err := p.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM categories").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := p.db.QueryContext(ctx, "SELECT * FROM categories LIMIT $1 OFFSET $2", pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, 0, err
		}
		categories = append(categories, category)
	}
	return categories, total, nil
}

func (p *CategoryRepositoryImpl) GetCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
	var category domain.Category
	err := p.db.QueryRowContext(ctx, "SELECT * FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (p *CategoryRepositoryImpl) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	query := `INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id`

	err := p.db.QueryRowContext(
		ctx,
		query,
		category.Name,
		category.Description,
	).Scan(&category.ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (p *CategoryRepositoryImpl) UpdateCategory(ctx context.Context, id int, category *domain.Category) (*domain.Category, error) {
	query := `UPDATE categories SET name = $1, description = $2 WHERE id = $3 RETURNING id`

	err := p.db.QueryRowContext(
		ctx,
		query,
		category.Name,
		category.Description,
		id,
	).Scan(&category.ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (p *CategoryRepositoryImpl) DeleteCategory(ctx context.Context, id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
