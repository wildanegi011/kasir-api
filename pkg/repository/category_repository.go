package repository

import (
	"database/sql"
	"kasir-api/pkg/domain"
)

type CategoryRepository interface {
	GetCategories() ([]domain.Category, error)
	GetCategoryByID(id int) (*domain.Category, error)
	CreateCategory(category *domain.Category) (*domain.Category, error)
	UpdateCategory(id int, category *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error
}

type CategoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (p *CategoryRepositoryImpl) GetCategories() ([]domain.Category, error) {
	rows, err := p.db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (p *CategoryRepositoryImpl) GetCategoryByID(id int) (*domain.Category, error) {
	var category domain.Category
	err := p.db.QueryRow("SELECT * FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (p *CategoryRepositoryImpl) CreateCategory(category *domain.Category) (*domain.Category, error) {
	query := `
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
		RETURNING id
	`

	err := p.db.QueryRow(
		query,
		category.Name,
		category.Description,
	).Scan(&category.ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (p *CategoryRepositoryImpl) UpdateCategory(id int, category *domain.Category) (*domain.Category, error) {
	query := `
		UPDATE categories SET name = $1, description = $2 WHERE id = $3
		RETURNING id
	`

	err := p.db.QueryRow(
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

func (p *CategoryRepositoryImpl) DeleteCategory(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
