package home

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
)

func GetAdminHomeAcademicData(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT 
    COUNT(CASE WHEN placement_fa >= 50.00 THEN roll_no END) AS placement_above_50,
    COUNT(CASE WHEN placement_fa <= 50.00 THEN roll_no END) AS placement_below_50,
    COUNT(CASE WHEN arrears >= 1 THEN roll_no END) AS backlog_students,
    COUNT(CASE WHEN cgpa >= 8.00 THEN roll_no END) AS cgpa_above_8,
    COUNT(CASE WHEN cgpa <= 8.00 THEN roll_no END) AS cgpa_below_8
FROM student_academic_data;

	`

	var placement_above_50, placement_below_50, backlog_students, cgpa_above_8, cgpa_below_8 int

	err := config.Database.QueryRow(query).Scan(&placement_above_50, &placement_below_50, &backlog_students, &cgpa_above_8,&cgpa_below_8)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}

	response := struct {
		PlaementFA_above50    int `json:"placement_above_50"`
		PlaementFA_below50 int `json:"placement_below_50"`
		BacklogCount       int `json:"backlog_students"`
		CGPA_above_8     int `json:"cgpa_above_8"`
		CGPA_below_8     int `json:"cgpa_below_8"`
	}{
		PlaementFA_above50:    placement_above_50,
		PlaementFA_below50: placement_below_50,
		BacklogCount:       backlog_students,
		CGPA_above_8:     cgpa_above_8,
		CGPA_below_8:     cgpa_below_8,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}
