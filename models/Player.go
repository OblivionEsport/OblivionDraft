package models

type Player struct {
	ChampionID int `json:"championId"`
	Team       int `json:"team"`
	SummonerID int `json:"summonerId"`
}
