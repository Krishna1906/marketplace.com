package payment

import (
	"github.com/gorilla/mux"
)

func RegisterPaymentRoutes(r *mux.Router) {
	r.HandleFunc("", GetPaymentMethods).Methods("GET")
	r.HandleFunc("/card", SaveCardPayment).Methods("POST")

}
