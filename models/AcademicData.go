package models

type AcademicData struct {
	Name        string  `json:"name"`
	RollNo      string  `json:"roll_no"`
	PlacementFa float64 `json:"placement_fa"`
	AcademicFa  float64 `json:"academic_fa"`
	Cgpa		float64 `json:"cgpa"`
	Sem1        float64 `json:"sem1"`
	Sem2        float64 `json:"sem2"`
	Sem3        float64 `json:"sem3"`
	Sem4        float64 `json:"sem4"`
	Sem5        float64 `json:"sem5"`
	Sem6        float64 `json:"sem6"`
	Sem7        float64 `json:"sem7"`
	Sem8        float64 `json:"sem8"`
	Arrears     int     `json:"arrears"`
}
