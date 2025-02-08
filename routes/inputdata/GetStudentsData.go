package inputdata

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/models"
)


func GetStudentsProfileData(w http.ResponseWriter, r *http.Request) {

	query := `SELECT roll_no, NAME, dob, gender, phone_no, regulation, batch, department, YEAR, mentor_id, mentor_name, aadhaar_no, pan_no, college_email, personal_email, parent_mobile, leetcode_username, github_username FROM users_profile WHERE ROLE="3"`

	rows, err := config.Database.Query(query)
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students [] models.Usersprofile
	for rows.Next() {
		var student models.Usersprofile
		if err := rows.Scan(
			&student.Rollno, &student.Name, &student.DOB, &student.Gender, &student.Phoneno, &student.Regulation,
			&student.Batch, &student.Department, &student.Year, &student.Mentor_id, &student.Mentor_name,
			&student.Aadhar_no, &student.Pan_no, &student.College_email, &student.Personal_email,
			&student.Parent_mobile_no, &student.Leetcode_username, &student.Github_username);
			err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	// Convert to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
