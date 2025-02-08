package models

type FullStackCardData struct {
	FullStackPoints int    `json:"full_stack_point"`
	FullStackRank   int    `json:"full_stack_rank"`
	Currentlevel    int    `json:"current_level"`
	AssignedStacks  string `json:"assigned_stack"`
}

type FullStackProjectDetails struct {
	Projectid     string `json:"project_id"`
	AssignedStack string `json:"assigned_stack"`
	Projecttitle  string `json:"project_title"`
	Projectdesc   string `json:"project_description"`
	Assignedby    string `json:"assigned_by"`
	Reviewername  string `json:"reviewer_name"`
	Revieweremail string `json:"reviewer_email"`
	Rvevenuename  string `json:"rve_name"`
	Reviewdate    string `json:"date"`
	Time          string `json:"time"`
}
