package models

type PersonalDetails struct {
	Rollno            string `json:"rollno"`
	Name              string `json:"name"`
	DOB               string `json:"dob"`
	Batch             string `json:"batch"`
	Phoneno           int    `json:"phone_no"`
	College_email     string `json:"college_email"`
	Personal_email    string `json:"personal_email"`
	Parent_mobile_no  int    `json:"parent_mobile_no"`
	Regulation        string `json:"regulation"`
	Department        string `json:"department"`
	Mentor_id         string `json:"mentor_id"`
	Mentor_name       string `json:"mentor_name"`
	Aadhar_no         int    `json:"aadhar_no"`
	Pan_no            string `json:"pan_no"`
}
