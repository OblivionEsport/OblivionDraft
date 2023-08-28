package models

import (
	"encoding/json"
	"fmt"
)

type Session struct {
	Actions         [][]Action `json:"actions"`
	Bans            Bans       `json:"bans"`
	GameID          int        `json:"gameId"`
	MyTeam          []Player   `json:"myTeam"`
	TheirTeam       []Player   `json:"theirTeam"`
	Counter         int        `json:"counter"`
	RecoveryCounter int        `json:"recoveryCounter"`
	Timer           Timer      `json:"timer"`
}

// make a string representation of the session
func (s *Session) String() string {
	return fmt.Sprintf("GameID: %d\nMyTeam: %v\nTheirTeam: %v\nBans: %v", s.GameID, s.MyTeam, s.TheirTeam, s.Bans)
}

// to JSON
func (s *Session) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
