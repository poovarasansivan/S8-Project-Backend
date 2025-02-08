package home

import (
	"server/config"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAdminHomeData(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT 
			COUNT(CASE WHEN p.placement_status = 'Placed' AND p.placement_type = 'IT' THEN up.roll_no END) AS placed_students,
			COUNT(CASE WHEN p.placement_status = 'Not Placed' AND p.placement_type = 'IT' THEN up.roll_no END) AS not_placed_students,
			COUNT(CASE WHEN p.placement_status = 'NA' AND p.placement_type = 'NIP' THEN up.roll_no END) AS nip_students,
			COUNT(up.roll_no) AS total_students
		FROM users_profile up
		LEFT JOIN placement p ON p.roll_no = up.roll_no
		WHERE up.role = '3';
	`

	var placedStudents, notPlacedStudents, nipStudents, totalStudents int

	err := config.Database.QueryRow(query).Scan(&placedStudents, &notPlacedStudents, &nipStudents, &totalStudents)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}

	response := struct {
		PlacedStudents   int `json:"placed_students"`
		NotPlacedStudents int `json:"not_placed_students"`
		NIPStudents      int `json:"nip_students"`
		TotalStudents    int `json:"total_students"`
	}{
		PlacedStudents:   placedStudents,
		NotPlacedStudents: notPlacedStudents,
		NIPStudents:      nipStudents,
		TotalStudents:    totalStudents,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}
