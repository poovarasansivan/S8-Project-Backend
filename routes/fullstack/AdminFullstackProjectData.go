package fullstack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"github.com/gorilla/mux"
)

func GetIndividualStudentFullStackData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Roll_No := vars["roll_no"]

	query := `
		SELECT roll_no, full_stack_point, full_stack_rank, current_level, assigned_stack
		FROM student_full_stack_map
		WHERE roll_no = ?;
	`

	var rollNo, fullStackPoint, fullStackRank, currentLevel, assignedStack string

	err := config.Database.QueryRow(query, Roll_No).Scan(&rollNo, &fullStackPoint, &fullStackRank, &currentLevel, &assignedStack)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}

	response := struct {
		RollNo          string `json:"roll_no"`
		FullStackPoint  string `json:"full_stack_point"`
		FullStackRank   string `json:"full_stack_rank"`
		CurrentLevel    string `json:"current_level"`
		AssignedStack   string `json:"assigned_stack"`
	}{
		RollNo:         rollNo,
		FullStackPoint: fullStackPoint,
		FullStackRank:  fullStackRank,
		CurrentLevel:   currentLevel,
		AssignedStack:  assignedStack,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}


func GetIndividualStudentProjectDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Roll_No := vars["roll_no"]

	query := `
		SELECT sfm.roll_no, up.name, fsp.project_title, fsp.project_description, fsp.assigend_by
		FROM student_full_stack_map sfm
		INNER JOIN full_stack_projects fsp ON fsp.project_id = sfm.project_id
		INNER JOIN users_profile up ON up.roll_no = sfm.roll_no
		WHERE sfm.roll_no = ?;
	`
	var rollNo, name, projectTitle, projectDescription, assignedBy string

	err := config.Database.QueryRow(query, Roll_No).Scan(&rollNo, &name, &projectTitle, &projectDescription, &assignedBy)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}

	response := struct {
		RollNo           string `json:"roll_no"`
		Name             string `json:"name"`
		ProjectTitle     string `json:"project_title"`
		ProjectDescription string `json:"project_description"`
		AssignedBy       string `json:"assigned_by"`
	}{
		RollNo:           rollNo,
		Name:             name,
		ProjectTitle:     projectTitle,
		ProjectDescription: projectDescription,
		AssignedBy:       assignedBy,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}


func GetIndividualStudentReviewData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Roll_No := vars["roll_no"]

	query := `
		SELECT up.name, fsr.roll_no, fsr.reviewer_name, fsr.reviewer_email, fsr.rev_department, fsr.rev_venue, fsr.date, fsr.time
		FROM users_profile up
		INNER JOIN full_stack_review fsr ON fsr.roll_no = up.roll_no
		WHERE fsr.roll_no = ?;
	`

	rows, err := config.Database.Query(query, Roll_No)
	if err != nil {
		// If an error occurs, return an error message
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}
	defer rows.Close()

	var reviews []struct {
		Name            string `json:"name"`
		RollNo          string `json:"roll_no"`
		ReviewerName    string `json:"reviewer_name"`
		ReviewerEmail   string `json:"reviewer_email"`
		RevDepartment   string `json:"rev_department"`
		RevVenue        string `json:"rev_venue"`
		Date            string `json:"date"`
		Time            string `json:"time"`
	}

	for rows.Next() {
		var review struct {
			Name            string `json:"name"`
			RollNo          string `json:"roll_no"`
			ReviewerName    string `json:"reviewer_name"`
			ReviewerEmail   string `json:"reviewer_email"`
			RevDepartment   string `json:"rev_department"`
			RevVenue        string `json:"rev_venue"`
			Date            string `json:"date"`
			Time            string `json:"time"`
		}

		err := rows.Scan(&review.Name, &review.RollNo, &review.ReviewerName, &review.ReviewerEmail, &review.RevDepartment, &review.RevVenue, &review.Date, &review.Time)
		if err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			fmt.Println("Row scan error:", err)
			return
		}

		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error fetching rows", http.StatusInternalServerError)
		fmt.Println("Rows error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Send the array of JSON objects
	err = json.NewEncoder(w).Encode(reviews)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}

