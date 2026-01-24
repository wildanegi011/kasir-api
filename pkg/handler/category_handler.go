package handler

import (
	"encoding/json"
	"errors"
	"kasir-api/pkg/dto"
	"kasir-api/pkg/service"
	"kasir-api/pkg/utils"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Mengambil semua data kategori
// @Tags categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} map[string]interface{}
// @Router /api/categories [get]
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	categories, total, err := h.categoryService.GetCategories(page, pageSize)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed to get categories")
		return
	}

	utils.SuccessResponse(
		w,
		http.StatusOK,
		"Categories found",
		categories,
		utils.WithPagination(total, page, pageSize),
	)
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Mengambil kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	category, err := h.categoryService.GetCategoryByID(idInt)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, utils.ErrCategoryNotFound.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Category found", category)
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Membuat kategori baru
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CategoryRequest true "Category Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/categories [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req dto.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ValidationErrorResponse(w, err)
		return
	}

	category := dto.CategoryReqToDomain(&req)

	if _, err := h.categoryService.CreateCategory(category); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create category")
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, "Category created successfully", category)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body dto.CategoryRequest true "Category Data"
// @Success 200 {object} map[string]interface{}
// @Router /api/categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	idInt, _ := strconv.Atoi(id)

	var req dto.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ValidationErrorResponse(w, err)
		return
	}

	category := dto.CategoryReqToDomain(&req)
	updatedCategory, err := h.categoryService.UpdateCategory(idInt, category)
	if err != nil {
		if errors.Is(err, utils.ErrCategoryNotFound) {
			utils.ErrorResponse(w, http.StatusNotFound, utils.ErrCategoryNotFound.Error())
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update category")
		return
	}

	updatedCategory.ID = idInt

	utils.SuccessResponse(w, http.StatusOK, "Category updated successfully", updatedCategory)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Menghapus kategori berdasarkan ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	if err := h.categoryService.DeleteCategory(idInt); err != nil {
		if errors.Is(err, utils.ErrCategoryNotFound) {
			utils.ErrorResponse(w, http.StatusNotFound, utils.ErrCategoryNotFound.Error())
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed to delete category")
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Category deleted successfully", nil)
}
