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
// @Success 200 {object} map[string]interface{}
// @Router /api/categories [get]
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "failed to get categories",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "successfully get categories",
		Data:    categories,
	})
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
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Invalid Request",
		})
		return
	}
	category, err := h.categoryService.GetCategoryByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "failed to get category",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "successfully get category",
		Data:    category,
	})
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
	w.Header().Set("Content-Type", "application/json")

	var req dto.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Name required",
		})
		return
	}

	category := dto.CategoryReqToDomain(&req)

	if _, err := h.categoryService.CreateCategory(category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Failed to create category",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "Category created successfully",
		Data:    category,
	})
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
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	idInt, _ := strconv.Atoi(id)

	var req dto.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Invalid request body",
		})
		return
	}

	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Name required",
		})
		return
	}

	category := dto.CategoryReqToDomain(&req)
	updatedCategory, err := h.categoryService.UpdateCategory(idInt, category)
	if err != nil {
		if errors.Is(err, utils.ErrCategoryNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.Response{
				Status:  false,
				Message: utils.ErrCategoryNotFound.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Failed to update category",
		})
		return
	}

	updatedCategory.ID = idInt

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "Category updated successfully",
		Data:    updatedCategory,
	})
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
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "Invalid Request",
		})
		return
	}
	if err := h.categoryService.DeleteCategory(idInt); err != nil {
		if errors.Is(err, utils.ErrCategoryNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.Response{
				Status:  false,
				Message: utils.ErrCategoryNotFound.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  false,
			Message: "failed to delete category",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Response{
		Status:  true,
		Message: "successfully delete category",
	})
}
