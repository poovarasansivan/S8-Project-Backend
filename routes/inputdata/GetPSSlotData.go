package inputdata

import (
	"encoding/json"
	"net/http"
	"server/config"
	"server/models"
)

func GetPsSlotData(w http.ResponseWriter, r *http.Request) {

	query := `SELECT pd.roll_no,up.name,pd.ps_category,pd.category_level,pd.slot_date,pd.level_status,pd.slot_missed,pd.negative_marks FROM ps_data pd INNER JOIN users_profile up ON up.roll_no=pd.roll_no`

	rows, err := config.Database.Query(query)
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.PSSlotData
	for rows.Next() {
		var student models.PSSlotData
		if err := rows.Scan(
			&student.Rollno,&student.Name, &student.PsCategory, &student.Category, &student.Slot, &student.LevelStatus, &student.SlotMissed, &student.Negativemarks); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	// Convert to JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
