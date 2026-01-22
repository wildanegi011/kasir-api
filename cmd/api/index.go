package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"kasir-api/pkg/config"
	appHandler "kasir-api/pkg/handler"
	"kasir-api/pkg/repository"
	"kasir-api/pkg/service"
	"kasir-api/pkg/utils"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	once sync.Once
)

// init routes sekali saja (cold start)
func initApp() {
	db, _ := config.InitDB()

	// ================= Product =================
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := appHandler.NewProductHandler(productService)

	http.HandleFunc("GET /api/products", productHandler.GetProducts)
	http.HandleFunc("GET /api/products/", productHandler.GetProductByID)
	http.HandleFunc("POST /api/products", productHandler.CreateProduct)
	http.HandleFunc("PUT /api/products/", productHandler.UpdateProduct)
	http.HandleFunc("DELETE /api/products/", productHandler.DeleteProduct)

	// ================= Category =================
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := appHandler.NewCategoryHandler(categoryService)

	http.HandleFunc("GET /api/categories", categoryHandler.GetCategories)
	http.HandleFunc("GET /api/categories/", categoryHandler.GetCategoryByID)
	http.HandleFunc("POST /api/categories", categoryHandler.CreateCategory)
	http.HandleFunc("PUT /api/categories/", categoryHandler.UpdateCategory)
	http.HandleFunc("DELETE /api/categories/", categoryHandler.DeleteCategory)

	// ================= Health =================
	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(utils.Response{
			Status:  true,
			Message: "server is running",
		})
	})

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initApp)
	http.DefaultServeMux.ServeHTTP(w, r)
}
