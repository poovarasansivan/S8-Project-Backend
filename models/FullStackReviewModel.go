package models

type FullStackReview struct {
	Rollno string `json:"roll_no"`
	Name string `json:"name"`
    ReviewerName string	`json:"reviewer_name"`
	ReviewerEmail string	`json:"reviewer_email"`
	ReviewerDepartment string	`json:"rev_department"`
	ReviewVenue string	`json:"rev_venue"`
	ReviewDate string	`json:"date"`
    ReviewTime string	`json:"time"`
}