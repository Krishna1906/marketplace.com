package banner

import (
	"marketplace/internal/database"
	"strconv"
)

// ✅ CREATE
func CreateBannerRepo(imageURL, startTime, endTime string) error {
	query := `
	INSERT INTO banners (image_url, start_time, end_time)
	VALUES ($1, $2, $3)
	`
	_, err := database.DB.Exec(query, imageURL, startTime, endTime)
	return err
}

// ✅ ADMIN GET ALL
func GetAllBannersRepo() ([]Banner, error) {
	query := `
	SELECT id, image_url, is_active, start_time, end_time, created_at
	FROM banners
	ORDER BY id DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []Banner

	for rows.Next() {
		var b Banner
		rows.Scan(&b.ID, &b.ImageURL, &b.IsActive, &b.StartTime, &b.EndTime, &b.CreatedAt)
		banners = append(banners, b)
	}

	return banners, nil
}

// ✅ USER GET (TIME FILTER 🔥)
func GetActiveBannersRepo() ([]Banner, error) {
	query := `
	SELECT id, image_url
	FROM banners
	WHERE is_active = true
	AND (start_time IS NULL OR start_time <= NOW())
	AND (end_time IS NULL OR end_time >= NOW())
	ORDER BY id DESC
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []Banner

	for rows.Next() {
		var b Banner
		rows.Scan(&b.ID, &b.ImageURL)
		banners = append(banners, b)
	}

	return banners, nil
}

func UpdateBannerRepoPartial(id string, startTime, endTime *string) error {

	query := "UPDATE banners SET "
	args := []interface{}{}
	i := 1

	if startTime != nil {
		query += "start_time = $" + strconv.Itoa(i)
		args = append(args, *startTime)
		i++
	}

	if endTime != nil {
		if len(args) > 0 {
			query += ", "
		}
		query += "end_time = $" + strconv.Itoa(i)
		args = append(args, *endTime)
		i++
	}

	query += " WHERE id = $" + strconv.Itoa(i)
	args = append(args, id)

	_, err := database.DB.Exec(query, args...)
	return err
}