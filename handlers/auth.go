package handlers

import (
	"server/config"
	"server/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("sunapana-suma-irukathada")

func Signin(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}
	var dbUser models.Users
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = config.Database.QueryRow("SELECT roll_no,NAME,college_email,department,YEAR,leetcode_username,github_username,ROLE FROM users_profile WHERE college_email=?", input.Email).Scan(
		&dbUser.Rollno, &dbUser.Name, &dbUser.College_email, &dbUser.Department, &dbUser.Year, &dbUser.Leetcode_username, &dbUser.Github_username, &dbUser.Role)

	if err == sql.ErrNoRows {
		http.Error(w, "Email not found", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": dbUser.Rollno,
		"role":   dbUser.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":      tokenString,
		"userID":     dbUser.Rollno,
		"username":   dbUser.Name,
		"role":       dbUser.Role,
		"department": dbUser.Department,
		"college_email": dbUser.College_email,
		"year": dbUser.Year,
		"leetcode_username": dbUser.Leetcode_username,
		"github_username": dbUser.Github_username,

	})
}
