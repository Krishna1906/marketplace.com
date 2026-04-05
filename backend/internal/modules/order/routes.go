package order

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterOrderRoutes(router *mux.Router) {
	router.HandleFunc("/create", CreateOrder).Methods("POST")
	router.HandleFunc("/payment", GetPaymentMethods).Methods("GET")
	// router.HandleFunc("/{order_id}", GetOrderDetails).Methods("GET")
	// router.HandleFunc("/user", GetUserOrders).Methods("GET")
	router.HandleFunc("/place", PlaceOrderFromCart).Methods("POST")
	router.HandleFunc("", GetOrders).Methods("GET")
	router.HandleFunc("/cancel/{id}", CancelOrder).Methods("PUT")
	router.HandleFunc("/track/{order_id}", TrackOrder).Methods("GET")
}

// GetPaymentMethods handles the request to retrieve payment methods
func GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment methods retrieved successfully"))
}
