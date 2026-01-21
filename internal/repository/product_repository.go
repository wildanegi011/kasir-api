package repository

import (
	"database/sql"
	"kasir-api/internal/model"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	GetProductByID(id int) (*model.Product, error)
	CreateProduct(product *model.Product) (*model.Product, error)
	UpdateProduct(id int, product *model.Product) (*model.Product, error)
	DeleteProduct(id int) error
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepositoryImpl(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (p *ProductRepositoryImpl) GetProducts() ([]model.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepositoryImpl) GetProductByID(id int) (*model.Product, error) {
	row := p.db.QueryRow("SELECT * from products WHERE id = ?", id)
	var product model.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepositoryImpl) CreateProduct(product *model.Product) (*model.Product, error) {
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

func (p *ProductRepositoryImpl) UpdateProduct(id int, product *model.Product) (*model.Product, error) {
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
