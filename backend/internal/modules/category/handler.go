package category

import (
	"encoding/json"
	"net/http"

	"marketplace/internal/database"
	"marketplace/internal/modules/product"
	"github.com/gorilla/mux"
)

// shared logic (PRIVATE)
func fetchActiveCategories(w http.ResponseWriter) {
	rows, err := database.DB.Query(`
		SELECT id, name, is_active, created_at
		FROM categories
		WHERE is_active = true
		ORDER BY name
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var c Category
		rows.Scan(
			&c.ID,
			&c.Name,
			&c.IsActive,
			&c.CreatedAt,
		)
		categories = append(categories, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// ---------------- USER (PUBLIC) ----------------
func GetUserCategories(w http.ResponseWriter, r *http.Request) {
	fetchActiveCategories(w)
}

// ---------------- SELLER ----------------
func GetSellerCategories(w http.ResponseWriter, r *http.Request) {
	fetchActiveCategories(w)
}

// ---------------- ADMIN ----------------

// POST /api/admin/categories
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var c Category

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := database.DB.QueryRow(`
		INSERT INTO categories (name, is_active)
		VALUES ($1, true)
		RETURNING id
	`, c.Name).Scan(&c.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

// GET /api/admin/categories
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, name, is_active, created_at
		FROM categories
		ORDER BY created_at DESC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		rows.Scan(&c.ID, &c.Name, &c.IsActive, &c.CreatedAt)
		categories = append(categories, c)
	}

	json.NewEncoder(w).Encode(categories)
}

// PUT /api/admin/categories/{id}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var c Category
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`
		UPDATE categories
		SET name = $1, is_active = $2
		WHERE id = $3
	`, c.Name, c.IsActive, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Category updated"))
}

// DELETE /api/admin/categories/{id}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := database.DB.Exec(`
		DELETE FROM categories WHERE id = $1
	`, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Category deleted"))
}

// ---------------- USER (PUBLIC) ----------------
// GET /api/categories/{id}
func GetCategoryProducts(w http.ResponseWriter, r *http.Request) {
	categoryID := mux.Vars(r)["id"]

	rows, err := database.DB.Query(`
		SELECT p.id, p.name, p.description, p.price, p.stock,
		       p.seller_id, p.category_id, c.name,
		       p.status, p.created_at
		FROM products p
		JOIN categories c ON c.id = p.category_id
		WHERE p.category_id = $1
		  AND p.status = 'APPROVED'
		ORDER BY p.created_at DESC
	`, categoryID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []product.Product

	for rows.Next() {
		var p product.Product
		rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.SellerID,
			&p.CategoryID,
			&p.Category,
			&p.Status,
			&p.CreatedAt,
		)

		p.Images, _ = product.GetProductImages(p.ID)
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
