package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kasir-api/pkg/config"
	"kasir-api/pkg/handler"
	"kasir-api/pkg/repository"
	"kasir-api/pkg/service"
	"kasir-api/pkg/utils"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Kasir API
// @version 1.0
// @host kasir-api-chi.vercel.app/
// @BasePath /

func main() {

	db, _ := config.InitDB()

	// =================== Product ===================================
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	http.HandleFunc("GET /api/products", productHandler.GetProducts)
	http.HandleFunc("GET /api/products/", productHandler.GetProductByID)
	http.HandleFunc("POST /api/products", productHandler.CreateProduct)
	http.HandleFunc("PUT /api/products/", productHandler.UpdateProduct)
	http.HandleFunc("DELETE /api/products/", productHandler.DeleteProduct)
	// =================================================================

	// =================== Category ===================================
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	http.HandleFunc("GET /api/categories", categoryHandler.GetCategories)
	http.HandleFunc("GET /api/categories/", categoryHandler.GetCategoryByID)
	http.HandleFunc("POST /api/categories", categoryHandler.CreateCategory)
	http.HandleFunc("PUT /api/categories/", categoryHandler.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/", categoryHandler.DeleteCategory)
	// =================================================================

	// =================== Health ===================================
	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.Response{
			Status:  true,
			Message: "server is running",
		})
	})

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Println("server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("gagal running server")
	}
}
