package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/produk"
	"net/http"
)

func main() {
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			produk.AmbilProdukByID(w, r)
		} else if r.Method == "PUT" {
			produk.UbahProdukByID(w, r)
		} else if r.Method == "DELETE" {
			produk.HapusProdukID(w, r)
		}
	})

	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk.DataProduk)
		} else if r.Method == "POST" {
			produk.TambahProdukBaru(w, r)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di localhost:8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
