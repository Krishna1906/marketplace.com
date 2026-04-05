package cart

func UpdateCartService(userID int64, productID string, qty int) (map[string]interface{}, error) {

	// ➖ If quantity <= 0 → remove
	if qty <= 0 {
		err := RemoveCartItemRepo(userID, productID)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"message": "Item removed",
		}, nil
	}

	// ➕ / ➖ update
	newQty, err := UpdateCartRepo(userID, productID, qty)
	if err != nil {
		return nil, err
	}

	// ❌ If becomes 0 → delete
	if newQty <= 0 {
		_ = RemoveCartItemRepo(userID, productID)
		return map[string]interface{}{
			"message": "Item removed",
		}, nil
	}

	return map[string]interface{}{
		"message":  "Cart updated",
		"quantity": newQty,
	}, nil
}