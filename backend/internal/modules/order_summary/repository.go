package ordersummary

import (
	"fmt"
	"marketplace/internal/database"
)

func GetCartItems(userID int64) ([]CartItem, float64, error) {
	query := `
	SELECT 
		p.id, p.name, p.price, c.quantity,
		(
			SELECT image_url 
			FROM product_images 
			WHERE product_id = p.id 
			LIMIT 1
		)
	FROM carts c
	JOIN products p ON p.id = c.product_id
	WHERE c.user_id = $1;
	`

	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []CartItem
	var total float64

	for rows.Next() {
		var item CartItem
		err := rows.Scan(&item.ProductID, &item.Name, &item.Price, &item.Quantity, &item.Image)
		if err != nil {
			return nil, 0, err
		}

		item.Total = item.Price * float64(item.Quantity)
		total += item.Total
		items = append(items, item)
	}

	return items, total, nil
}

func GetAddresses(userID int64) ([]Address, error) {
	query := `
	SELECT id, type, full_name, phone, address_line, city, state, pincode
	FROM addresses
	WHERE user_id = $1
	ORDER BY id DESC
	`

	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []Address

	for rows.Next() {
		var a Address
		err := rows.Scan(
			&a.ID,
			&a.Type,
			&a.FullName,
			&a.Phone,
			&a.AddressLine,
			&a.City,
			&a.State,
			&a.Pincode,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, a)
	}

	return addresses, nil
}

func CreateAddress(userID int64, a Address) error {
	query := `
	INSERT INTO addresses 
	(user_id, type, full_name, phone, address_line, city, state, pincode)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`

	_, err := database.DB.Exec(query,
		userID,
		a.Type,
		a.FullName,
		a.Phone,
		a.AddressLine,
		a.City,
		a.State,
		a.Pincode,
	)

	if err != nil {
		fmt.Println("DB ERROR (Create):", err)
	}

	return err
}

func UpdateAddress(userID, addressID int64, a Address) error {
	query := `
	UPDATE addresses SET
		type=$1,
		full_name=$2,
		phone=$3,
		address_line=$4,
		city=$5,
		state=$6,
		pincode=$7
	WHERE id=$8 AND user_id=$9
	`

	_, err := database.DB.Exec(query,
		a.Type,
		a.FullName,
		a.Phone,
		a.AddressLine,
		a.City,
		a.State,
		a.Pincode,
		addressID,
		userID,
	)

	if err != nil {
		fmt.Println("DB ERROR (Update):", err)
	}

	return err
}

func DeleteAddress(userID, addressID int64) error {
	query := `DELETE FROM addresses WHERE id=$1 AND user_id=$2`

	_, err := database.DB.Exec(query, addressID, userID)

	if err != nil {
		fmt.Println("DB ERROR (Delete):", err)
	}

	return err
}