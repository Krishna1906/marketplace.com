package admin

import (
	"net/http"
	"time"

	"marketplace/internal/database"
	"marketplace/internal/utils"
)

type AdminUserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	query := `
		SELECT id, name, email, role, created_at
		FROM users
		ORDER BY created_at DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []AdminUserResponse

	for rows.Next() {
		var u AdminUserResponse
		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Role,
			&u.CreatedAt,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	utils.JSON(w, http.StatusOK, users)
}
