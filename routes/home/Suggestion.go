package home

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/config"
)

type Suggestion struct {
	RollNo        string `json:"roll_no"`
	Comments      string `json:"comments"`
	ReadingStatus string `json:"reading_status"`
}

func InsertSuggestion(w http.ResponseWriter, r *http.Request) {

	var suggestion Suggestion
	err := json.NewDecoder(r.Body).Decode(&suggestion)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		log.Println("Error reading request body:", err)
		return
	}

	query := `INSERT INTO suggestion(roll_no, comments, reading_status, created_at, updated_at) 
			  VALUES (?, ?, "0", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	_, err = config.Database.Exec(query, suggestion.RollNo, suggestion.Comments)
	if err != nil {
		http.Error(w, "Unable to insert data into database", http.StatusInternalServerError)
		log.Println("Database query error:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Suggestion successfully inserted",
		"roll_no": suggestion.RollNo,
	}
	json.NewEncoder(w).Encode(response)
}


func GetIndividualSuggestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Roll_No := vars["roll_no"]
	query := `SELECT roll_no, comments, reading_status FROM suggestion WHERE roll_no=?`
	var suggestions []struct {
		RollNo        string `json:"roll_no"`
		Comments      string `json:"comments"`
		ReadingStatus string `json:"reading_status"`
	}
	rows, err := config.Database.Query(query, Roll_No)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var suggestion struct {
			RollNo        string `json:"roll_no"`
			Comments      string `json:"comments"`
			ReadingStatus string `json:"reading_status"`
		}
		err := rows.Scan(&suggestion.RollNo, &suggestion.Comments, &suggestion.ReadingStatus)
		if err != nil {
			http.Error(w, "Error reading data from database", http.StatusInternalServerError)
			fmt.Println("Error reading data from database:", err)
			return
		}
		suggestions = append(suggestions, suggestion)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "Error processing rows", http.StatusInternalServerError)
		fmt.Println("Error processing rows:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(suggestions)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}
}

func UpdateSuggestionStatus(w http.ResponseWriter, r *http.Request) {
	var suggestion Suggestion
	err := json.NewDecoder(r.Body).Decode(&suggestion)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		fmt.Println("Error reading request body:", err)
		return
	}

	query := `UPDATE suggestion SET reading_status=? WHERE roll_no=?`
	_, err = config.Database.Exec(query, suggestion.ReadingStatus, suggestion.RollNo)
	if err != nil {
		http.Error(w, "Unable to update data in the database", http.StatusInternalServerError)
		fmt.Println("Database update error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}