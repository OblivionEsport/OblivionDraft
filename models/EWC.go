package models

type GameState struct {
	Data  Data   `json:"data"`
	Event string `json:"event"`
}

type Data struct {
	Event     string               `json:"event"`
	Game      Game                 `json:"game"`
	HasGame   bool                 `json:"hasGame"`
	MatchGUID string               `json:"match_guid"`
	Players   map[string]PlayerEWC `json:"players"`
}

type Game struct {
	Arena            string    `json:"arena"`
	Ball             Ball      `json:"ball"`
	HasTarget        bool      `json:"hasTarget"`
	HasWinner        bool      `json:"hasWinner"`
	IsOT             bool      `json:"isOT"`
	IsReplay         bool      `json:"isReplay"`
	Target           string    `json:"target"`
	Teams            []TeamEWC `json:"teams"`
	TimeMilliseconds float64   `json:"time_milliseconds"`
	TimeSeconds      int       `json:"time_seconds"`
	Winner           string    `json:"winner"`
}

type Ball struct {
	Speed int `json:"speed"`
	Team  int `json:"team"`
}

type TeamEWC struct {
	ColorPrimary   string `json:"color_primary"`
	ColorSecondary string `json:"color_secondary"`
	Name           string `json:"name"`
	Score          int    `json:"score"`
}

type PlayerEWC struct {
	Assists        int      `json:"assists"`
	Attacker       string   `json:"attacker"`
	Boost          int      `json:"boost"`
	Cartouches     int      `json:"cartouches"`
	Demos          int      `json:"demos"`
	Goals          int      `json:"goals"`
	HasCar         bool     `json:"hasCar"`
	ID             string   `json:"id"`
	IsDead         bool     `json:"isDead"`
	IsPowersliding bool     `json:"isPowersliding"`
	IsSonic        bool     `json:"isSonic"`
	Location       Location `json:"location"`
	Name           string   `json:"name"`
	OnGround       bool     `json:"onGround"`
	OnWall         bool     `json:"onWall"`
	PrimaryID      string   `json:"primaryID"`
	Saves          int      `json:"saves"`
	Score          int      `json:"score"`
	Shortcut       int      `json:"shortcut"`
	Shots          int      `json:"shots"`
	Speed          int      `json:"speed"`
	Team           int      `json:"team"`
	Touches        int      `json:"touches"`
}

type Location struct {
	X     float64 `json:"X"`
	Y     float64 `json:"Y"`
	Z     float64 `json:"Z"`
	Pitch int     `json:"pitch"`
	Roll  int     `json:"roll"`
	Yaw   int     `json:"yaw"`
}
