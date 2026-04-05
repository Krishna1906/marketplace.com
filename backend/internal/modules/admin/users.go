package admin

import (
	"net/http"

	"marketplace/internal/database"
	"marketplace/internal/utils"
)

type UserResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

func FetchAllUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := database.DB.Query(`
		SELECT id, name, email, role, created_at
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []UserResponse

	for rows.Next() {
		var u UserResponse
		rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Role,
			&u.CreatedAt,
		)
		users = append(users, u)
	}

	utils.JSON(w, http.StatusOK, users)
}
