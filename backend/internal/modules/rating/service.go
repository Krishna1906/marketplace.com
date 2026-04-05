package rating

import "errors"

// ➕ Add / Update
func SaveRating(userID, productID int64, rating int) error {

	if rating < 1 || rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	// 🔥 CHECK DELIVERY
	canRate, err := CanUserRate(userID, productID)
	if err != nil {
		return err
	}

	if !canRate {
		return errors.New("you can rate only after delivery")
	}

	return AddOrUpdateRating(userID, productID, rating)
}

// 📊 Get Rating
func FetchRating(productID int64) (float64, int, error) {
	return GetProductRating(productID)
}