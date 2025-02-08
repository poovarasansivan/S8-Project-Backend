package technical

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"

	"github.com/gorilla/mux"
)

func GetCodingPlatformsUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Roll_No := vars["roll_no"]

	query := `SELECT leetcode_username,github_username FROM users_profile WHERE roll_no=?`
	var Leetcodeusername, Githubusername string
	err := config.Database.QueryRow(query, Roll_No).Scan(&Leetcodeusername, &Githubusername)
	if err != nil {
		http.Error(w, "Unable to fetch data from database", http.StatusInternalServerError)
		fmt.Println("Database query error:", err)
		return
	}

	response := struct {
		Leetcodename string `json:"leetcode_username"`
		Githubname   string `json:"github_username"`
	}{
		Leetcodename: Leetcodeusername,
		Githubname:   Githubusername,
	}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding data to JSON", http.StatusInternalServerError)
		fmt.Println("JSON encoding error:", err)
		return
	}

}
