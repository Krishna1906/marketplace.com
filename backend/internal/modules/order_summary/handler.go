package ordersummary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"marketplace/internal/middleware"
	"marketplace/internal/utils"

	"github.com/gorilla/mux"
)

// 🔹 GET
func GetOrderSummary(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "Unauthorized", 401)
		return
	}

	items, total, err := GetCartItems(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	addresses, err := GetAddresses(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.JSON(w, http.StatusOK, OrderSummaryResponse{
		Items:       items,
		TotalAmount: total,
		Addresses:   addresses,
	})
}

// 🔹 POST
func SaveOrderAddress(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "Unauthorized", 401)
		return
	}

	var addr Address
	if err := json.NewDecoder(r.Body).Decode(&addr); err != nil {
		http.Error(w, "Invalid body", 400)
		return
	}

	fmt.Println("Incoming Address:", addr)

	err := CreateAddress(userID, addr)
	if err != nil {
		http.Error(w, "Failed to save address", 500)
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"message": "Address added successfully",
	})
}

// 🔹 PUT
func UpdateOrderAddress(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "Unauthorized", 401)
		return
	}

	params := mux.Vars(r)
	addressID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	var addr Address
	if err := json.NewDecoder(r.Body).Decode(&addr); err != nil {
		http.Error(w, "Invalid body", 400)
		return
	}

	err = UpdateAddress(userID, addressID, addr)
	if err != nil {
		http.Error(w, "Failed to update address", 500)
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"message": "Address updated successfully",
	})
}

// 🔹 DELETE
func DeleteOrderAddress(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		http.Error(w, "Unauthorized", 401)
		return
	}

	params := mux.Vars(r)
	addressID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	err = DeleteAddress(userID, addressID)
	if err != nil {
		http.Error(w, "Failed to delete address", 500)
		return
	}

	utils.JSON(w, http.StatusOK, map[string]string{
		"message": "Address deleted successfully",
	})
}