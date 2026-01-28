package handlers

import (
	"encoding/json"
	"errors"
	"kasir-api/internal/dto"
	service "kasir-api/internal/services"
	"kasir-api/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Mengambil semua data produk
// @Tags products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} map[string]interface{}
// @Router /api/products [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	products, total, err := h.productService.GetProducts(r.Context(), page, pageSize)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(
		w,
		http.StatusOK,
		"Products found",
		products,
		utils.WithPagination(total, page, pageSize),
	)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Mengambil produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/products/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	product, err := h.productService.GetProductByID(r.Context(), idInt)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed to get product")
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Product found", product)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Membuat produk baru
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.ProductRequest true "Product Data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ValidationErrorResponse(w, err)
		return
	}

	product := dto.ProductReqToDomain(&req)

	if _, err := h.productService.CreateProduct(r.Context(), product); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, "Product created successfully", product)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body dto.ProductRequest true "Product Data"
// @Success 200 {object} map[string]interface{}
// @Router /api/products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/products/")
	idInt, _ := strconv.Atoi(id)

	var req dto.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ValidationErrorResponse(w, err)
		return
	}

	product := dto.ProductReqToDomain(&req)
	updatedProduct, err := h.productService.UpdateProduct(r.Context(), idInt, product)
	if err != nil {
		if errors.Is(err, utils.ErrProductNotFound) {
			utils.ErrorResponse(w, http.StatusNotFound, utils.ErrProductNotFound.Error())
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	updatedProduct.ID = idInt

	utils.SuccessResponse(w, http.StatusOK, "Product updated successfully", updatedProduct)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Menghapus produk berdasarkan ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/products/")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	if err := h.productService.DeleteProduct(r.Context(), idInt); err != nil {
		if errors.Is(err, utils.ErrProductNotFound) {
			utils.ErrorResponse(w, http.StatusNotFound, utils.ErrProductNotFound.Error())
			return
		}
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed to delete product")
		return
	}

	utils.SuccessResponse(w, http.StatusOK, "Product deleted successfully", nil)
}
