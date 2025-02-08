package home

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
)

type StudentProfile struct {
	RollNo          string `json:"roll_no"`
	Name            string `json:"name"`
	DOB             string `json:"dob"`
	Gender          string `json:"gender"`
	PhoneNo         string `json:"phone_no"`
	Regulation      string `json:"regulation"`
	Batch           string `json:"batch"`
	Department      string `json:"department"`
	Year            string `json:"year"`
	MentorID        string `json:"mentor_id"`
	MentorName      string `json:"mentor_name"`
	AadhaarNo       string `json:"aadhaar_no"`
	PanNo           string `json:"pan_no"`
	CollegeEmail    string `json:"college_email"`
	PersonalEmail   string `json:"personal_email"`
	ParentMobile    string `json:"parent_mobile"`
	LeetCodeUsername string `json:"leetcode_username"`
	GitHubUsername  string `json:"github_username"`
	Role            string `json:"role"`
}

func AddStudentsProfileData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var students []StudentProfile
	if err := json.NewDecoder(r.Body).Decode(&students); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	
	tx, err := config.Database.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	stmt, err := tx.Prepare("INSERT INTO users_profile (roll_no, name, dob, gender, phone_no, regulation, batch, department, year, mentor_id, mentor_name, aadhaar_no, pan_no, college_email, personal_email, parent_mobile, leetcode_username, github_username, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	for _, student := range students {
		_, err := stmt.Exec(student.RollNo, student.Name, student.DOB, student.Gender, student.PhoneNo, student.Regulation, student.Batch, student.Department, student.Year, student.MentorID, student.MentorName, student.AadhaarNo, student.PanNo, student.CollegeEmail, student.PersonalEmail, student.ParentMobile, student.LeetCodeUsername, student.GitHubUsername, student.Role)
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Failed to insert student data: %v", err), http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Student profiles inserted successfully"))
}