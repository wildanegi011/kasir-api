package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
	"kasir-api/internal/utils"
)

type ProductService interface {
	GetProducts() ([]model.Product, error)
	GetProductByID(id int) (*model.Product, error)
	CreateProduct(product *model.Product) (*model.Product, error)
	UpdateProduct(id int, product *model.Product) (*model.Product, error)
	DeleteProduct(id int) error
}

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (s *ProductServiceImpl) GetProducts() ([]model.Product, error) {
	return s.productRepository.GetProducts()
}

func (s *ProductServiceImpl) GetProductByID(id int) (*model.Product, error) {
	return s.productRepository.GetProductByID(id)
}

func (s *ProductServiceImpl) CreateProduct(product *model.Product) (*model.Product, error) {
	return s.productRepository.CreateProduct(product)
}

func (s *ProductServiceImpl) UpdateProduct(id int, product *model.Product) (*model.Product, error) {
	_, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return nil, utils.ErrProductNotFound
	}
	return s.productRepository.UpdateProduct(id, product)
}

func (s *ProductServiceImpl) DeleteProduct(id int) error {
	return s.productRepository.DeleteProduct(id)
}
