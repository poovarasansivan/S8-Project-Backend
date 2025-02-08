package inputdata

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/models"
)

func GetFullStackProjectData(w http.ResponseWriter, r *http.Request) {

	query := `SELECT sfm.roll_no,up.name,sfm.full_stack_point,sfm.full_stack_rank,sfm.current_level,sfm.assigned_stack,sfm.project_id,fsp.project_title,fsp.project_description,fsp.assigend_by FROM student_full_stack_map sfm INNER JOIN full_stack_projects fsp ON fsp.project_id=sfm.project_id INNER JOIN users_profile up ON up.roll_no=sfm.roll_no`
	rows, err := config.Database.Query(query)
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.FullStackProjectData
	for rows.Next() {
		var student models.FullStackProjectData
		if err := rows.Scan(
			&student.Rollno,&student.Name, &student.Fullstackpoint, &student.FullStackRank, &student.CurrentLevel, &student.AssignedStack, &student.ProjectId, &student.Title, &student.Description, &student.Assignedby); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	// Convert to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
