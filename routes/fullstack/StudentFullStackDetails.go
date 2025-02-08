package fullstack

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"server/config"
	"server/models"
)

func GetStudentFullStackCardData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT full_stack_point,full_stack_rank,current_level,assigned_stack FROM student_full_stack_map WHERE roll_no=?`

	var fullstack models.FullStackCardData
	err := config.Database.QueryRow(query, rollNo).Scan(
		&fullstack.FullStackPoints,
		&fullstack.FullStackRank,
		&fullstack.Currentlevel,
		&fullstack.AssignedStacks,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(fullstack)

}

func GetFullStackCurrentlevel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT current_level FROM student_full_stack_map WHERE roll_no=?`

	var currentlevel int
	err := config.Database.QueryRow(query, rollNo).Scan(&currentlevel)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		CurrentLevel int `json:"current_level"`
	}{
		CurrentLevel: currentlevel,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}

func GetProjectDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT fsm.project_id,fsm.assigned_stack,fsp.project_title,fsp.project_description,fsp.assigend_by,fsr.reviewer_name,fsr.reviewer_email,fsr.rev_venue,fsr.date,fsr.time FROM student_full_stack_map fsm INNER JOIN full_stack_projects fsp ON fsp.project_id = fsm.project_id INNER JOIN full_stack_review fsr 
    ON fsr.roll_no=fsm.roll_no WHERE fsm.roll_no=?`
	var fullstack models.FullStackProjectDetails
	err := config.Database.QueryRow(query, rollNo).Scan(
		&fullstack.Projectid,
		&fullstack.AssignedStack,
		&fullstack.Projecttitle,
		&fullstack.Projectdesc,
		&fullstack.Assignedby,
		&fullstack.Reviewername,
		&fullstack.Revieweremail,
		&fullstack.Rvevenuename,
		&fullstack.Reviewdate,
		&fullstack.Time,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(fullstack)

}
