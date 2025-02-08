package technical

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
)

func GetStudentRollno(w http.ResponseWriter, r *http.Request) {

	query := `
		SELECT roll_no FROM users_profile WHERE ROLE="3"
	`
	var rollNos [] string

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
