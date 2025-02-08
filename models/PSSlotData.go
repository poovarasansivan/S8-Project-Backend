package models

type PSSlotData struct {
	Rollno        string `json:"roll_no"`
	Name          string `json:"name"`
	PsCategory    string `json:"ps_category"`
	Category      string `json:"category_level"`
	Slot          string `json:"slot_date"`
	LevelStatus   string `json:"level_status"`
	SlotMissed    string `json:"slot_missed"`
	Negativemarks string `json:"negative_marks"`
}
