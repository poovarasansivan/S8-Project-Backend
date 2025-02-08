package inputdata

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/models"
)

func GetAcademicData(w http.ResponseWriter, r *http.Request) {

	query := `SELECT up.name,ad.roll_no, ad.placement_fa, ad.academic_fa, ad.cgpa,ad.sem1,ad.sem2,ad.sem3,ad.sem4,ad.sem5,ad.sem6,ad.sem7,ad.sem8,ad.arrears FROM student_academic_data ad INNER JOIN users_profile up ON 
up.roll_no = ad.roll_no  `

	rows, err := config.Database.Query(query)
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.AcademicData
	for rows.Next() {
		var student models.AcademicData
		if err := rows.Scan(&student.Name,
			&student.RollNo, &student.PlacementFa, &student.AcademicFa, &student.Cgpa, &student.Sem1, &student.Sem2, &student.Sem3, &student.Sem4, &student.Sem5, &student.Sem6, &student.Sem7, &student.Sem8, &student.Arrears); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	// Convert to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
