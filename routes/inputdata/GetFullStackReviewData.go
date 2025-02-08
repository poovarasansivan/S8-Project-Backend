package inputdata

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/models"
)

func GetFullStackReviewData(w http.ResponseWriter, r *http.Request) {

	query := `SELECT fsr.roll_no,up.name,fsr.reviewer_name,fsr.reviewer_email,fsr.rev_department,fsr.rev_venue,fsr.date,fsr.time FROM full_stack_review fsr INNER JOIN users_profile up ON up.roll_no = fsr.roll_no`

	rows, err := config.Database.Query(query)
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.FullStackReview
	for rows.Next() {
		var student models.FullStackReview
		if err := rows.Scan(
			&student.Rollno,&student.Name, &student.ReviewerName, &student.ReviewerEmail, &student.ReviewerDepartment, &student.ReviewVenue, &student.ReviewDate, &student.ReviewTime); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	// Convert to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
