package models

type StudentEducationProfile struct {
	SSLCSchoolName    string  `json:"sslc_school_name"`
	SSLCBoard         string  `json:"sslc_board"`
	SSLCYearCompleted string  `json:"sslc_year_completed"`
	SSLCPercentage    float64 `json:"sslc_percentage"`
	HSCSchoolName     string  `json:"hsc_school_name"`
	HSCBoard          string  `json:"hsc_board"`
	HSCYearCompleted  string  `json:"hsc_year_completed"`
	HSCPercentage     float64 `json:"hsc_percentage"`
	DegreeName        string  `json:"degree_name"`
	Department        string  `json:"department"`
	CollegeName       string  `json:"clg_name"`
	YearofGraduation  string  `json:"year_graduation"`
	CGPA              float64 `json:"cgpa"`
}
