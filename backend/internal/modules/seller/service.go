package seller

import "errors"

func ApplyForSeller(userID int64, shopName, gst string) error {
	seller := &Seller{
		UserID:    userID,
		ShopName: shopName,
		GSTNumber: gst,
		Status:    "PENDING",
	}
	return CreateSeller(seller)
}

func EnsureSellerApproved(userID int64) error {
	seller, err := GetSellerByUserID(userID)
	if err != nil {
		return errors.New("seller not found")
	}
	if seller.Status != "APPROVED" {
		return errors.New("seller not approved")
	}
	return nil
}
