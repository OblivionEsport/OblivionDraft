package models

type Action struct {
	ActorCellId  int    `json:"actorCellId"`
	ChampionId   int    `json:"championId"`
	Completed    bool   `json:"completed"`
	Id           int    `json:"id"`
	IsAllyAction bool   `json:"isAllyAction"`
	IsInProgress bool   `json:"isInProgress"`
	Type         string `json:"type"`
}
