package admin

import (
	"errors"

	"marketplace/internal/modules/auth"
	"marketplace/internal/modules/seller"
)

func ApproveSeller(userID int64) error {
	if userID == 0 {
		return errors.New("invalid user id")
	}

	err := seller.ApproveSellerByUserID(userID)
	if err != nil {
		return err
	}

	return auth.UpdateUserRole(userID, "SELLER")
}
