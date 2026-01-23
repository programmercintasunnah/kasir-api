package categories

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var DataCategories = []Category{
	{ID: 1, Name: "Mie Instan", Description: "Aneka mie instan seperti Indomie, Mie Sedaap"},
	{ID: 2, Name: "Minuman Kemasan", Description: "Air mineral, teh botol, minuman ringan"},
	{ID: 3, Name: "Snack", Description: "Cemilan seperti biskuit, wafer, keripik"},
	{ID: 4, Name: "Bahan Dapur", Description: "Kecap, saus, gula, garam"},
}

/* GET /categories/{id} */
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, c := range DataCategories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

/* POST /categories */
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	category.ID = len(DataCategories) + 1
	DataCategories = append(DataCategories, category)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

/* PUT /categories/{id} */
func UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updated Category
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i := range DataCategories {
		if DataCategories[i].ID == id {
			updated.ID = id
			DataCategories[i] = updated

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

/* DELETE /categories/{id} */
func DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, c := range DataCategories {
		if c.ID == id {
			DataCategories = append(DataCategories[:i], DataCategories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "category deleted",
			})
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
