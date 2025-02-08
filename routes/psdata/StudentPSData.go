package psdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/models"

	"github.com/gorilla/mux"
)

func GetStudentPSCardData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT ps_category, COUNT(ps_category) 
	          FROM ps_data 
	          WHERE roll_no = ? 
	          GROUP BY ps_category;`

	rows, err := config.Database.Query(query, rollNo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Ensure rows are closed after function execution

	// Slice to store multiple rows
	var psdataList []models.PSCardData

	for rows.Next() {
		var psdata models.PSCardData
		err := rows.Scan(&psdata.PsCategory, &psdata.LevelsCompleted)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning data: %v", err), http.StatusInternalServerError)
			return
		}
		psdataList = append(psdataList, psdata)
	}

	// Check for errors from iteration
	if err = rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error iterating over data: %v", err), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send JSON response
	json.NewEncoder(w).Encode(psdataList)
}

type PSData struct {
	TotalLevelCompleted int `json:"total_level_completed"`
}

func GetStudentPSData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT COUNT(ps_category) AS totallevelcompleted FROM ps_data WHERE roll_no=?`

	var psdata PSData // Define the struct instance

	err := config.Database.QueryRow(query, rollNo).Scan(&psdata.TotalLevelCompleted)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send JSON response
	json.NewEncoder(w).Encode(psdata)
}

func GetSudentPsPointHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `SELECT ps_category,category_level, 
       DATE_FORMAT(slot_date, '%Y-%m-%d') AS completion_date, 
       DATE_FORMAT(slot_date, '%M') AS completion_month
       FROM ps_data 
       WHERE roll_no = ? 
       ORDER BY created_at DESC 
       LIMIT 7;`

	rows, err := config.Database.Query(query, rollNo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close() 
	var psdataList []models.PointsHistory

	for rows.Next() {
		var psdata models.PointsHistory
		err := rows.Scan(&psdata.PSCategroy, &psdata.CategoryLevel,&psdata.CompletionDate, &psdata.CompletedMonth)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning data: %v", err), http.StatusInternalServerError)
			return
		}
		psdataList = append(psdataList, psdata)
	}

	// Check for errors from iteration
	if err = rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error iterating over data: %v", err), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send JSON response
	json.NewEncoder(w).Encode(psdataList)
}


func GetPSlevelsData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNo := vars["roll_no"]

	query := `
		SELECT pd.ps_category, COUNT(pd.ps_category) AS levels_completed, pod.total_levels
		FROM ps_ovr_data pod
		INNER JOIN ps_data pd ON pd.ps_category = pod.category
		WHERE roll_no = ? 
		GROUP BY ps_category`

	var psdata [] models.PsLeveldetails

	rows, err := config.Database.Query(query, rollNo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var levelData models.PsLeveldetails
		err := rows.Scan(&levelData.PsCategory, &levelData.LevelsCompleted, &levelData.TotalLevels)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning data: %v", err), http.StatusInternalServerError)
			return
		}
		psdata = append(psdata, levelData)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error processing rows: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(psdata)
}
