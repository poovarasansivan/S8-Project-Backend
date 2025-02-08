package models

type PsLeveldetails struct {
	PsCategory      string `json:"ps_category"`
	LevelsCompleted int    `json:"levels_completed"`
	TotalLevels    int    `json:"total_levels"`
}
