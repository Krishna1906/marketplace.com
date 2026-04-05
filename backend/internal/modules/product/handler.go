package product

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"marketplace/internal/database"
	"marketplace/internal/middleware"
	"marketplace/internal/utils"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(middleware.UserIDKey)

	// PUBLIC OR NORMAL USER → show marketplace products
	if userIDVal == nil {
		getPublicProducts(w, r)
		return
	}

	role := r.Context().Value(middleware.RoleKey)
	if role != "SELLER" && role != "ADMIN" {
		getPublicProducts(w, r)
		return
	}

	// SELLER → show own products
	sellerID := userIDVal.(int64)

	query := `
		SELECT p.id, p.name, p.description, p.price, p.stock,
       p.seller_id,
       p.category_id, c.name,
       p.status, p.created_at
		FROM products p
		JOIN categories c ON c.id = p.category_id
		WHERE p.seller_id = $1
		ORDER BY p.created_at DESC

	`
	rows, err := database.DB.Query(query, sellerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(
				&p.ID,
				&p.Name,
				&p.Description,
				&p.Price,
				&p.Stock,
				&p.SellerID,
				&p.CategoryID,
				&p.CategoryName,
				&p.Status,
				&p.CreatedAt,
			)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p.Images, _ = GetProductImages(p.ID)

		products = append(products, p)
	}

	utils.JSON(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	// ✅ Safe context extraction
	userIDVal := r.Context().Value(middleware.UserIDKey)
	sellerID, ok := userIDVal.(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO products 
		(name, description, price, stock, seller_id, category_id, status)
		VALUES ($1, $2, $3, $4, $5, $6, 'PENDING')
		RETURNING id
		`

	if product.CategoryID == 0 {
		http.Error(w, "category_id is required", http.StatusBadRequest)
		return
	}

	err = database.DB.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		sellerID,
		product.CategoryID,
	).Scan(&product.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product created successfully"))
}

func getPublicProducts(w http.ResponseWriter, r *http.Request) {

	categoryID := r.URL.Query().Get("category")

	baseQuery := `
	SELECT p.id, p.name, p.description, p.price, p.stock,
	       p.category_id, c.name, p.status, p.created_at
	FROM products p
	JOIN categories c ON c.id = p.category_id
	WHERE p.status = 'APPROVED'
	`

	var rows *sql.Rows
	var err error

	if categoryID != "" {
		baseQuery += " AND p.category_id = $1"
		rows, err = database.DB.Query(baseQuery, categoryID)
	} else {
		rows, err = database.DB.Query(baseQuery)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.CategoryID,
			&p.Category,
			&p.Status,
			&p.CreatedAt,
		)
		p.Images, _ = GetProductImages(p.ID)
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}

func UploadProductImages(w http.ResponseWriter, r *http.Request) {

	sellerID := r.Context().Value(middleware.UserIDKey).(int64)
	productID := mux.Vars(r)["id"]

	// 🔐 Verify seller owns the product
	var ownerID int64
	err := database.DB.QueryRow(
		"SELECT seller_id FROM products WHERE id = $1",
		productID,
	).Scan(&ownerID)

	if err != nil || ownerID != sellerID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse multipart form (10MB max)
	err = r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		http.Error(w, "No images uploaded", http.StatusBadRequest)
		return
	}

	for _, fileHeader := range files {

		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		filename := fmt.Sprintf(
			"%d_%s",
			time.Now().UnixNano(),
			fileHeader.Filename,
		)

		filePath := filepath.Join("uploads/products", filename)
		out, _ := os.Create(filePath)
		defer out.Close()

		io.Copy(out, file)

		imageURL := "/uploads/products/" + filename

		_, err = database.DB.Exec(
			"INSERT INTO product_images (product_id, image_url) VALUES ($1, $2)",
			productID,
			imageURL,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Images uploaded successfully"))
}

func GetProductImages(productID int64) ([]string, error) {
	rows, err := database.DB.Query(
		"SELECT image_url FROM product_images WHERE product_id = $1",
		productID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []string
	for rows.Next() {
		var url string
		rows.Scan(&url)
		images = append(images, url)
	}
	return images, nil
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {

	productID := mux.Vars(r)["id"]

	query := `
	SELECT p.id, p.name, p.description, p.price, p.stock,
	       p.seller_id, p.category_id, c.name,
	       p.status, p.created_at
	FROM products p
	JOIN categories c ON c.id = p.category_id
	WHERE p.id = $1 AND p.status = 'APPROVED'
	`

	var p Product
	err := database.DB.QueryRow(query, productID).Scan(
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

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	p.Images, _ = GetProductImages(p.ID)

	utils.JSON(w, http.StatusOK, p)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	userIDVal := r.Context().Value(middleware.UserIDKey)
	sellerID, ok := userIDVal.(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	productIDStr := mux.Vars(r)["id"]

	productID, err := strconv.ParseInt(productIDStr, 10, 64)
if err != nil {
	http.Error(w, "Invalid product ID", http.StatusBadRequest)
	return
}

	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 🔎 verify product belongs to seller
	var ownerID int64
		err = database.DB.QueryRow(
			"SELECT seller_id FROM products WHERE id = $1",
			productID,
		).Scan(&ownerID)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if ownerID != sellerID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if req.CategoryID == 0 {
	http.Error(w, "category_id is required", http.StatusBadRequest)
	return
}

	// ✅ update without changing status
	_, err = database.DB.Exec(`
	UPDATE products
	SET name = $1,
	    description = $2,
	    price = $3,
	    stock = $4,
	    category_id = $5
	WHERE id = $6 AND seller_id = $7
`,
	req.Name,
	req.Description,
	req.Price,
	req.Stock,
	req.CategoryID,
	productID,
	sellerID,
)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"message": "Product updated successfully",
	})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	sellerID := r.Context().Value(middleware.UserIDKey).(int64)
	productID := mux.Vars(r)["id"]

	_, err := database.DB.Exec(`
		UPDATE products
		SET status = 'DELETED'
		WHERE id = $1 AND seller_id = $2
	`, productID, sellerID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"message": "Product deleted successfully",
	})
}

func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {

	categoryID := mux.Vars(r)["id"]

	rows, err := database.DB.Query(`
		SELECT 
			p.id,
			p.name,
			p.description,
			p.price,
			p.stock,
			p.seller_id,
			p.category_id,
			p.status,
			p.created_at,
			COALESCE(
				ARRAY_AGG(pi.image_url) 
				FILTER (WHERE pi.image_url IS NOT NULL),
				'{}'
			) AS images
		FROM products p
		LEFT JOIN product_images pi ON p.id = pi.product_id
		WHERE p.category_id = $1
		  AND p.status = 'APPROVED'
		GROUP BY p.id
		ORDER BY p.created_at DESC
	`, categoryID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.SellerID,
			&p.CategoryID,
			&p.Status,
			&p.CreatedAt,
			&p.Images,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	utils.JSON(w, http.StatusOK, products)
}
