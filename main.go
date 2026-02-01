package main

import (
	"encoding/json"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	services "kasir-api/sevices"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	// ✅ DB_PORT KHUSUS HTTP SERVER (DARI RAILWAY)
	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatal("DB_PORT is not set (Railway Web Service required)")
	}

	// ✅ DB CONNECTION
	dbConn := viper.GetString("DB_CONN")
	if dbConn == "" {
		log.Fatal("DB_CONN is not set")
	}

	// Setup database
	db, err := database.InitDB(dbConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/api/category", categoryHandler.HandleCategory)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

	addr := ":" + port
	log.Println("Server running on", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
