package service

import (
	"kasir-api/internal/domain"
	"kasir-api/internal/repository"
	"kasir-api/internal/utils"
)

type ProductService interface {
	GetProducts() ([]domain.Product, error)
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

func (s *ProductServiceImpl) GetProducts() ([]domain.Product, error) {
	return s.productRepository.GetProducts()
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
