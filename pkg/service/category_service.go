package service

import (
	"kasir-api/pkg/domain"
	"kasir-api/pkg/repository"
	"kasir-api/pkg/utils"
)

type CategoryService interface {
	GetCategories() ([]domain.Category, error)
	GetCategoryByID(id int) (*domain.Category, error)
	CreateCategory(category *domain.Category) (*domain.Category, error)
	UpdateCategory(id int, category *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error
}

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{categoryRepository: categoryRepository}
}

func (s *CategoryServiceImpl) GetCategories() ([]domain.Category, error) {
	return s.categoryRepository.GetCategories()
}

func (s *CategoryServiceImpl) GetCategoryByID(id int) (*domain.Category, error) {
	return s.categoryRepository.GetCategoryByID(id)
}

func (s *CategoryServiceImpl) CreateCategory(category *domain.Category) (*domain.Category, error) {
	return s.categoryRepository.CreateCategory(category)
}

func (s *CategoryServiceImpl) UpdateCategory(id int, category *domain.Category) (*domain.Category, error) {
	_, err := s.categoryRepository.GetCategoryByID(id)
	if err != nil {
		return nil, utils.ErrCategoryNotFound
	}
	return s.categoryRepository.UpdateCategory(id, category)
}

func (s *CategoryServiceImpl) DeleteCategory(id int) error {
	_, err := s.categoryRepository.GetCategoryByID(id)
	if err != nil {
		return utils.ErrCategoryNotFound
	}
	return s.categoryRepository.DeleteCategory(id)
}
