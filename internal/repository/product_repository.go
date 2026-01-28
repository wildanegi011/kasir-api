package repository

import (
	"context"
	"database/sql"
	"kasir-api/internal/domain"
)

type ProductRepository interface {
	GetProducts(ctx context.Context, page int, pageSize int) ([]domain.Product, int, error)
	GetProductByID(ctx context.Context, id int) (*domain.Product, error)
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id int, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id int) error
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (p *ProductRepositoryImpl) GetProducts(ctx context.Context, page int, pageSize int) ([]domain.Product, int, error) {
	var total int
	err := p.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM products").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := p.db.QueryContext(ctx, "SELECT * FROM products LIMIT $1 OFFSET $2", pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CategoryID); err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}
	return products, total, nil
}

func (p *ProductRepositoryImpl) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
	var product domain.Product
	err := p.db.QueryRowContext(ctx, "SELECT * FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepositoryImpl) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	query := `INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id`

	err := p.db.QueryRowContext(
		ctx,
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

func (p *ProductRepositoryImpl) UpdateProduct(ctx context.Context, id int, product *domain.Product) (*domain.Product, error) {
	query := `UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4 RETURNING id`

	err := p.db.QueryRowContext(
		ctx,
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

func (p *ProductRepositoryImpl) DeleteProduct(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
