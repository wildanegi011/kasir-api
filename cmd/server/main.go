package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"kasir-api/internal/config"
	"kasir-api/internal/database"
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/service"
	"kasir-api/internal/utils"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Kasir API
// @version 1.0
// @host kasir-api-production-1c80.up.railway.app
// @BasePath /
func main() {
	// load config
	cfg := config.GetConfig()

	// connect to database
	db, closeDB, err := database.NewPostgres(&cfg.Database)
	if err != nil {
		panic("failed connect to database")
	}
	defer closeDB()

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

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.App.Port // local/dev
	}

	addr := ":" + port

	fmt.Println("server running di", addr)
	if err = http.ListenAndServe(addr, nil); err != nil {
		panic("failed running server")
	}
}
