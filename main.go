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

// ubah Config
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set by Railway")
	}

	config := Config{
		Port:   port,
		DBConn: viper.GetString("DB_CONN"),
	}

	if config.Port == "" {
		config.Port = "8088"
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
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

	// Setup routes
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/api/category", categoryHandler.HandleCategory)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

	addr := "0.0.0.0:" + config.Port
	log.Println("Server running on", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
