package models

type StudentCardData struct {
	PlacementStatus string `json:"placement_status"`
	FullStackPoints int    `json:"full_stack_point"`
	FullStackRank   int    `json:"full_stack_rank"`
}
