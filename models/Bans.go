package models

type Bans struct {
	MyTeamBans    []int `json:"myTeamBans"`
	NumBans       int   `json:"numBans"`
	TheirTeamBans []int `json:"theirTeamBans"`
}
