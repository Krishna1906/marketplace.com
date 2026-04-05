package banner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/mux"
)

// ✅ UPLOAD MULTIPLE BANNERS WITH TIME
func UploadBanner(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	startTime := r.FormValue("start_time")
	endTime := r.FormValue("end_time")

	files := r.MultipartForm.File["images"]

	for _, fileHeader := range files {

		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
		path := "uploads/banners/" + filename

		out, err := os.Create(path)
		if err != nil {
			continue
		}
		defer out.Close()

		_, err = out.ReadFrom(file)
		if err != nil {
			continue
		}

		imageURL := "/uploads/banners/" + filename

		UploadBannerService(imageURL, startTime, endTime)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Banners uploaded successfully"))
}

// ✅ ADMIN GET
func GetAdminBanners(w http.ResponseWriter, r *http.Request) {
	data, err := GetAllBannersService()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(data)
}

// ✅ USER GET
func GetUserBanners(w http.ResponseWriter, r *http.Request) {
	data, err := GetUserBannersService()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(data)
}


// ✅ UPDATE BANNER TIME
func UpdateBanner(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	type Req struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	var req Req

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}

	layout := "02-Jan-2006 15:04:05"

	var startStr *string
	var endStr *string

	// ✅ OPTIONAL START TIME
	if req.StartTime != "" {
		start, err := time.Parse(layout, req.StartTime)
		if err != nil {
			http.Error(w, "Invalid start_time format", 400)
			return
		}
		s := start.Format("2006-01-02 15:04:05")
		startStr = &s
	}

	// ✅ OPTIONAL END TIME
	if req.EndTime != "" {
		end, err := time.Parse(layout, req.EndTime)
		if err != nil {
			http.Error(w, "Invalid end_time format", 400)
			return
		}
		e := end.Format("2006-01-02 15:04:05")
		endStr = &e
	}

	err := UpdateBannerServicePartial(id, startStr, endStr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Banner updated successfully",
	})
}