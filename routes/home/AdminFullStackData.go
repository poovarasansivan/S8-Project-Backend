package home

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
)

func GetAdminHomeFullStackData(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT current_level, COUNT(*) AS student_count
		FROM student_full_stack_map
		GROUP BY current_level
		ORDER BY current_level;
	`

	rows, err := config.Database.Query(query)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}
	defer rows.Close()

	// Define a struct to hold individual level counts
	type LevelCount struct {
		CurrentLevel int `json:"current_level"`
		StudentCount int `json:"student_count"`
	}

	var levelCounts []LevelCount

	// Iterate through the results and store them in a slice
	for rows.Next() {
		var levelCount LevelCount
		if err := rows.Scan(&levelCount.CurrentLevel, &levelCount.StudentCount); err != nil {
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			fmt.Println("Row scan error:", err)
			return
		}
		levelCounts = append(levelCounts, levelCount)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		fmt.Println("Row iteration error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(levelCounts)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}
