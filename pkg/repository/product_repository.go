package repository

import (
	"database/sql"
	"kasir-api/pkg/domain"
)

type ProductRepository interface {
	GetProducts() ([]domain.Product, error)
	GetProductByID(id int) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(id int, product *domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (p *ProductRepositoryImpl) GetProducts() ([]domain.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepositoryImpl) GetProductByID(id int) (*domain.Product, error) {
	var product domain.Product
	err := p.db.QueryRow("SELECT * FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepositoryImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (name, price, stock)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := p.db.QueryRow(
		query,
		product.Name,
		product.Price,
		product.Stock,
	).Scan(&product.ID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepositoryImpl) UpdateProduct(id int, product *domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4
		RETURNING id
	`

	err := p.db.QueryRow(
		query,
		product.Name,
		product.Price,
		product.Stock,
		id,
	).Scan(&product.ID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepositoryImpl) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
