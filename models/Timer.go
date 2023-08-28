package models

type Timer struct {
	AdjustedTimeLeftInPhase int    `json:"adjustedTimeLeftInPhase"`
	Phase                   string `json:"phase"`
	TotalTimeInPhase        int    `json:"totalTimeInPhase"`
}
