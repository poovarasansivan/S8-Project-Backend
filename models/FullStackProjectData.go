package models

type FullStackProjectData struct {
	Rollno         string `json:"roll_no"`
	Name           string `json:"name"`
	Fullstackpoint int    `json:"full_stack_point"`
	FullStackRank  int    `json:"full_stack_rank"`
	CurrentLevel   int    `json:"current_level"`
	AssignedStack  string `json:"assigned_stack"`
	ProjectId      string `json:"project_id"`
	Title          string `json:"project_title"`
	Description    string `json:"project_description"`
	Assignedby     string `json:"assigned_by"`
}
