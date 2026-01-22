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
	row := p.db.QueryRow("SELECT * from products WHERE id = ?", id)
	var product domain.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepositoryImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	result, err := p.db.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)",
		product.Name, product.Price, product.Stock)
	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted product
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Set the ID of the product and return it
	product.ID = int(id)
	return product, nil
}

func (p *ProductRepositoryImpl) UpdateProduct(id int, product *domain.Product) (*domain.Product, error) {
	_, err := p.db.Exec("UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ?",
		product.Name, product.Price, product.Stock, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductRepositoryImpl) DeleteProduct(id int) error {
	_, err := p.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
