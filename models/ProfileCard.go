package models

type ProfileCard struct {
	Name       string `json:"name"`
	Batch      string `json:"batch"`
	Year       string `json:"year"`
	Department string `json:"department"`
	CGPA       float64 `json:"cgpa"`
}
