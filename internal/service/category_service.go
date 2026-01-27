package service

import (
	"kasir-api/internal/domain"
	"kasir-api/internal/repository"
	"kasir-api/internal/utils"
)

type CategoryService interface {
	GetCategories(page int, pageSize int) ([]domain.Category, int, error)
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

func (s *CategoryServiceImpl) GetCategories(page int, pageSize int) ([]domain.Category, int, error) {
	categories, total, err := s.categoryRepository.GetCategories(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	if categories == nil {
		categories = []domain.Category{}
	}

	return categories, total, nil
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
