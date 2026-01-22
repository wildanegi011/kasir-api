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
	row := p.db.QueryRow("SELECT * from categories WHERE id = ?", id)
	var category domain.Category
	if err := row.Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return nil, err
	}
	return &category, nil
}

func (p *CategoryRepositoryImpl) CreateCategory(category *domain.Category) (*domain.Category, error) {
	result, err := p.db.Exec("INSERT INTO categories (name, description) VALUES (?, ?)",
		category.Name, category.Description)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted product
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Set the ID of the product and return it
	category.ID = int(id)
	return category, nil
}

func (p *CategoryRepositoryImpl) UpdateCategory(id int, category *domain.Category) (*domain.Category, error) {
	_, err := p.db.Exec("UPDATE categories SET name = ?, description = ? WHERE id = ?",
		category.Name, category.Description, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (p *CategoryRepositoryImpl) DeleteCategory(id int) error {
	_, err := p.db.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
