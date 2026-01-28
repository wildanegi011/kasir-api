package service

import (
	"context"
	"kasir-api/internal/domain"
	"kasir-api/internal/repository"
	"kasir-api/internal/utils"
)

type CategoryService interface {
	GetCategories(ctx context.Context, page int, pageSize int) ([]domain.Category, int, error)
	GetCategoryByID(ctx context.Context, id int) (*domain.Category, error)
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	UpdateCategory(ctx context.Context, id int, category *domain.Category) (*domain.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{categoryRepository: categoryRepository}
}

func (s *CategoryServiceImpl) GetCategories(ctx context.Context, page int, pageSize int) ([]domain.Category, int, error) {
	categories, total, err := s.categoryRepository.GetCategories(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	if categories == nil {
		categories = []domain.Category{}
	}

	return categories, total, nil
}

func (s *CategoryServiceImpl) GetCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
	return s.categoryRepository.GetCategoryByID(ctx, id)
}

func (s *CategoryServiceImpl) CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return s.categoryRepository.CreateCategory(ctx, category)
}

func (s *CategoryServiceImpl) UpdateCategory(ctx context.Context, id int, category *domain.Category) (*domain.Category, error) {
	_, err := s.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, utils.ErrCategoryNotFound
	}
	return s.categoryRepository.UpdateCategory(ctx, id, category)
}

func (s *CategoryServiceImpl) DeleteCategory(ctx context.Context, id int) error {
	_, err := s.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		return utils.ErrCategoryNotFound
	}
	return s.categoryRepository.DeleteCategory(ctx, id)
}
