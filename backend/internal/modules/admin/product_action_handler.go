package admin

import (
	"encoding/json"
	"net/http"
)

type ProductActionRequest struct {
	ProductID int64  `json:"product_id"`
	Action    string `json:"action"` // APPROVED / REJECTED
}

func ProductActionHandler(w http.ResponseWriter, r *http.Request) {
	var req ProductActionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.ProductID == 0 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = UpdateProductStatus(req.ProductID, req.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Product status updated"))
}
