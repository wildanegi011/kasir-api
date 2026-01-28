package service

import (
	"context"
	"kasir-api/internal/domain"
	"kasir-api/internal/repository"
	"kasir-api/internal/utils"
)

type ProductService interface {
	GetProducts(ctx context.Context, page int, pageSize int) ([]domain.Product, int, error)
	GetProductByID(ctx context.Context, id int) (*domain.Product, error)
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id int, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id int) error
}

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (s *ProductServiceImpl) GetProducts(ctx context.Context, page int, pageSize int) ([]domain.Product, int, error) {
	products, total, err := s.productRepository.GetProducts(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	if products == nil {
		products = []domain.Product{}
	}

	return products, total, nil
}

func (s *ProductServiceImpl) GetProductByID(ctx context.Context, id int) (*domain.Product, error) {
	return s.productRepository.GetProductByID(ctx, id)
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	return s.productRepository.CreateProduct(ctx, product)
}

func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, id int, product *domain.Product) (*domain.Product, error) {
	_, err := s.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return nil, utils.ErrProductNotFound
	}
	return s.productRepository.UpdateProduct(ctx, id, product)
}

func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, id int) error {
	_, err := s.productRepository.GetProductByID(ctx, id)
	if err != nil {
		return utils.ErrProductNotFound
	}
	return s.productRepository.DeleteProduct(ctx, id)
}
