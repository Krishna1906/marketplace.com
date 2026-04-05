package seller

import "marketplace/internal/database"

func CreateSeller(s *Seller) error {
	query := `
		INSERT INTO sellers (user_id, shop_name, gst_number)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	return database.DB.QueryRow(
		query,
		s.UserID,
		s.ShopName,
		s.GSTNumber,
	).Scan(&s.ID)
}

func UpdateSellerStatus(id int64, status string) error {
	query := `UPDATE sellers SET status=$1 WHERE id=$2`
	_, err := database.DB.Exec(query, status, id)
	return err
}

func GetSellerByUserID(userID int64) (*Seller, error) {
	query := `
		SELECT id, user_id, shop_name, gst_number, status
		FROM sellers WHERE user_id=$1
	`
	row := database.DB.QueryRow(query, userID)

	s := &Seller{}
	err := row.Scan(&s.ID, &s.UserID, &s.ShopName, &s.GSTNumber, &s.Status)
	return s, err
}

func ApproveSellerByUserID(userID int64) error {
	query := `
		UPDATE sellers
		SET status = 'APPROVED'
		WHERE user_id = $1
	`
	_, err := database.DB.Exec(query, userID)
	return err
}

