package service

import (
	"kasir-api/pkg/domain"
	"kasir-api/pkg/repository"
	"kasir-api/pkg/utils"
)

type ProductService interface {
	GetProducts(page int, pageSize int) ([]domain.Product, int, error)
	GetProductByID(id int) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(id int, product *domain.Product) (*domain.Product, error)
	DeleteProduct(id int) error
}

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (s *ProductServiceImpl) GetProducts(page int, pageSize int) ([]domain.Product, int, error) {
	products, total, err := s.productRepository.GetProducts(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (s *ProductServiceImpl) GetProductByID(id int) (*domain.Product, error) {
	return s.productRepository.GetProductByID(id)
}

func (s *ProductServiceImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return s.productRepository.CreateProduct(product)
}

func (s *ProductServiceImpl) UpdateProduct(id int, product *domain.Product) (*domain.Product, error) {
	_, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return nil, utils.ErrProductNotFound
	}
	return s.productRepository.UpdateProduct(id, product)
}

func (s *ProductServiceImpl) DeleteProduct(id int) error {
	_, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return utils.ErrProductNotFound
	}
	return s.productRepository.DeleteProduct(id)
}
