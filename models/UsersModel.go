package models

type Users struct {
	Rollno            string `json:"rollno"`
	Name              string `json:"name"`
	DOB               string `json:"dob"`
	Gender            string `json:"gender"`
	Phoneno           int    `json:"phone_no"`
	Regulation        string `json:"regulation"`
	Batch             string `json:"batch"`
	Department        string `json:"department"`
	Year              string `json:"year"`
	Mentor_id         string `json:"mentor_id"`
	Mentor_name       string `json:"mentor_name"`
	Aadhar_no         int    `json:"aadhar_no"`
	Pan_no            string `json:"pan_no"`
	College_email     string `json:"college_email"`
	Personal_email    string `json:"personal_email"`
	Parent_mobile_no  int    `json:"parent_mobile_no"`
	Leetcode_username string `json:"leetcode_username"`
	Github_username   string `json:"github_username"`
	Role              string `json:"role"`
	Created_at        string `json:"created_at"`
	Updated_at        string `json:"updated_at"`
}


type Usersprofile struct {
	Rollno            string `json:"rollno"`
	Name              string `json:"name"`
	DOB               string `json:"dob"`
	Gender            string `json:"gender"`
	Phoneno           int    `json:"phone_no"`
	Regulation        string `json:"regulation"`
	Batch             string `json:"batch"`
	Department        string `json:"department"`
	Year              string `json:"year"`
	Mentor_id         string `json:"mentor_id"`
	Mentor_name       string `json:"mentor_name"`
	Aadhar_no         int    `json:"aadhar_no"`
	Pan_no            string `json:"pan_no"`
	College_email     string `json:"college_email"`
	Personal_email    string `json:"personal_email"`
	Parent_mobile_no  int    `json:"parent_mobile_no"`
	Leetcode_username string `json:"leetcode_username"`
	Github_username   string `json:"github_username"`
}