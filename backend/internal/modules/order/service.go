package order

import (
	"errors"
)

func CalculateTotal(price float64, qty int) float64 {
	return price * float64(qty)
}

func ValidatePayment(method string) error {
	if method == "" {
		return errors.New("payment_type is required")
	}
	if method != "COD" && method != "CARD" {
		return errors.New("invalid payment type")
	}
	return nil
}