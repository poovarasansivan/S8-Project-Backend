package psdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"

	"github.com/gorilla/mux"
)

func GetAdminPsCardData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Roll_No := vars["roll_no"]

	query := `
		SELECT 
  COUNT(category_level) AS levels_completed,
  COUNT(CASE WHEN negative_marks = 'yes' THEN 1 END) AS count_negative_marks
FROM ps_data
WHERE roll_no = ?
GROUP BY roll_no;

	`

	var levels_completed, count_negative_marks string

	err := config.Database.QueryRow(query, Roll_No).Scan(&levels_completed, &count_negative_marks)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}

	response := struct {
		LevelsCompleted string `json:"levels_completed"`
		Negativemarks   string `json:"count_negative_marks"`
	}{
		LevelsCompleted: levels_completed,
		Negativemarks:   count_negative_marks,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}



func GetPSDataRollNo(w http.ResponseWriter, r *http.Request) {
	query := `
			SELECT roll_no FROM ps_data GROUP BY roll_no
		`
	var rollNos []string

	rows, err := config.Database.Query(query)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var roll_no string
		if err := rows.Scan(&roll_no); err != nil {
			http.Error(w, "Error reading data from database", http.StatusInternalServerError)
			fmt.Println("Error reading data from database:", err)
			return
		}
		rollNos = append(rollNos, roll_no)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error fetching data from database", http.StatusInternalServerError)
		fmt.Println("Row iteration error:", err)
		return
	}

	response := struct {
		OverallStudentRollNos []string `json:"roll_no"`
	}{
		OverallStudentRollNos: rollNos,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}

func GetIndividualStudentPSSlotData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	// Check if the roll_no parameter is missing
	if rollNo == "" {
		http.Error(w, "Missing roll_no parameter", http.StatusBadRequest)
		return
	}

	// SQL query to fetch the data
	query := `
		SELECT 
			ps_category AS category, 
			category_level AS level, 
			slot_date AS slot_bookedtime,
			level_status, 
			slot_missed, 
			negative_marks 
		FROM ps_data 
		WHERE roll_no = ?;
	`

	// Execute the query
	rows, err := config.Database.Query(query, rollNo)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}
	defer rows.Close()

	type Review struct {
		Category       string `json:"category"`
		Level          string `json:"level"`
		SlotBookedTime string `json:"slot_bookedtime"`
		LevelStatus    string `json:"level_status"`
		SlotMissed     string `json:"slot_missed"`
		NegativeMarks  string `json:"negative_marks"`
	}

	var reviews []Review

	for rows.Next() {
		var review Review
		err := rows.Scan(&review.Category, &review.Level, &review.SlotBookedTime, &review.LevelStatus, &review.SlotMissed, &review.NegativeMarks)
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

	if len(reviews) == 0 {
		reviews = []Review{}
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(reviews)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
	}
}


func GetFacultyViewPSLevelCompletion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	if rollNo == "" {
		http.Error(w, "Missing roll_no parameter", http.StatusBadRequest)
		return
	}

	// SQL Query
	query := `
		SELECT 
			psd.ps_category, 
			pso.monthly_level, 
			COUNT(psd.category_level) AS completed_levels 
		FROM ps_data psd 
		INNER JOIN ps_ovr_data pso 
			ON pso.category = psd.ps_category 
		WHERE psd.roll_no = ? 
		GROUP BY psd.ps_category, pso.monthly_level;
	`

	rows, err := config.Database.Query(query, rollNo)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}
	defer rows.Close()

	type FacultyReview struct {
		Category        string `json:"category"`
		MonthlyLevel    string `json:"monthly_level"`
		CompletedLevels int    `json:"completed_levels"`
	}

	var facultyReviews []FacultyReview

	for rows.Next() {
		var review FacultyReview
		err := rows.Scan(&review.Category, &review.MonthlyLevel, &review.CompletedLevels)
		if err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			fmt.Println("Row scan error:", err)
			return
		}
		facultyReviews = append(facultyReviews, review)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error fetching rows", http.StatusInternalServerError)
		fmt.Println("Rows error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(facultyReviews)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
	}
}


func GetStudentAcademicDataForFacultyView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	if rollNo == "" {
		http.Error(w, "Missing roll_no parameter", http.StatusBadRequest)
		return
	}

	// SQL Query
	query := `
		SELECT 
			ad.roll_no, 
			up.name, 
			up.year, 
			up.department, 
			ad.cgpa 
		FROM users_profile up 
		INNER JOIN student_academic_data ad 
			ON ad.roll_no = up.roll_no 
		WHERE up.roll_no = ?;
	`

	type AcademicData struct {
		RollNo     string  `json:"roll_no"`
		Name       string  `json:"name"`
		Year       string  `json:"year"`
		Department string  `json:"department"`
		CGPA       float64 `json:"cgpa"`
	}

	var academicData AcademicData
	err := config.Database.QueryRow(query, rollNo).Scan(
		&academicData.RollNo,
		&academicData.Name,
		&academicData.Year,
		&academicData.Department,
		&academicData.CGPA,
	)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(academicData)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
	}
}
