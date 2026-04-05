package auth

import (
	"database/sql"
	"errors"

	"marketplace/internal/database"
)

func CreateUser(user *User) error {
	query := `
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return database.DB.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	).Scan(&user.ID)
}

func GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT id, name, email, password, role
		FROM users
		WHERE email = $1
	`

	row := database.DB.QueryRow(query, email)

	user := &User{}
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return user, err
}
func UpdateUserRole(userID int64, role string) error {
	query := `UPDATE users SET role=$1 WHERE id=$2`
	_, err := database.DB.Exec(query, role, userID)
	return err
}
