package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/models"

	"github.com/gorilla/mux"
)

func GetStudentProfileCardData(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	// Define the query
	query := `SELECT up.name, up.batch, up.year, up.department, sad.cgpa 
	          FROM users_profile up 
	          INNER JOIN student_academic_data sad ON sad.roll_no = up.roll_no 
	          WHERE up.roll_no = ?`

	var profile models.ProfileCard

	// Execute the query
	err := config.Database.QueryRow(query, rollNo).Scan(
		&profile.Name,
		&profile.Batch,
		&profile.Year,
		&profile.Department,
		&profile.CGPA,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send JSON response
	json.NewEncoder(w).Encode(profile)
}

func GetStudentCardData(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT tp.placement_status,fsm.full_stack_point,fsm.full_stack_rank FROM student_full_stack_map fsm INNER JOIN placement tp ON tp.roll_no = fsm.roll_no WHERE fsm.roll_no=?`

	var profile models.StudentCardData

	err := config.Database.QueryRow(query, rollNo).Scan(
		&profile.PlacementStatus,
		&profile.FullStackPoints,
		&profile.FullStackRank,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(profile)
}

func GetStudentPlacementDetails(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT placement_offer_1,offer1_lpa, placement_offer_2, offer2_lpa FROM placement WHERE roll_no=?`
	var profile models.Placementdetails

	err := config.Database.QueryRow(query, rollNo).Scan(
		&profile.Placementoffer1,
		&profile.Offer1_lpa,
		&profile.Placementoffer2,
		&profile.Offer2_lpa,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(profile)
}

func GetStudentSGPAScore(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT sem1,sem2,sem3,sem4,sem5,sem6,sem7,sem8 FROM student_academic_data WHERE roll_no=?`
	var profile models.StudentSGPA

	err := config.Database.QueryRow(query, rollNo).Scan(
		&profile.Sem1,
		&profile.Sem2,
		&profile.Sem3,
		&profile.Sem4,
		&profile.Sem5,
		&profile.Sem6,
		&profile.Sem7,
		&profile.Sem8,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(profile)
}

func GetStudentPersonalDetails(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT roll_no,NAME,dob,batch,phone_no,college_email,personal_email,parent_mobile,regulation,department,mentor_id,mentor_name,aadhaar_no,pan_no FROM users_profile WHERE roll_no=?`
	var profile models.PersonalDetails

	err := config.Database.QueryRow(query, rollNo).Scan(
		&profile.Rollno,
		&profile.Name,
		&profile.DOB,
		&profile.Batch,
		&profile.Phoneno,
		&profile.College_email,
		&profile.Personal_email,
		&profile.Parent_mobile_no,
		&profile.Regulation,
		&profile.Department,
		&profile.Mentor_id,
		&profile.Mentor_name,
		&profile.Aadhar_no,
		&profile.Pan_no,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(profile)
}

func GetStudentEducationProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT se.sslc_school_name,se.sslc_board,se.sslc_year_completed,se.sslc_percentage,se.hsc_school_name,se.hsc_board,se.hsc_year_completed,se.hsc_percentage,se.degree_name,se.department,se.year_graduation,se.clg_name,ad.cgpa
    FROM student_education se INNER JOIN student_academic_data ad ON ad.roll_no= se.roll_no WHERE se.roll_no=?`

	var profile models.StudentEducationProfile

	err := config.Database.QueryRow(query, rollNo).Scan(
		&profile.SSLCSchoolName,
		&profile.SSLCBoard,
		&profile.SSLCYearCompleted,
		&profile.SSLCPercentage,
		&profile.HSCSchoolName,
		&profile.HSCBoard,
		&profile.HSCYearCompleted,
		&profile.HSCPercentage,
		&profile.DegreeName,
		&profile.Department,
		&profile.YearofGraduation,
		&profile.CollegeName,
		&profile.CGPA,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(profile)
}
