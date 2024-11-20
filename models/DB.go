package models

type DBTournament struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Start     string `json:"start"`
	End       string `json:"end"`
	NameID    string `json:"name_id"`
	CreatedAt string `json:"created_at"`
}

type DBTeam struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Tag          string `json:"tag"`
	LogoUrl      string `json:"logo_url"`
	TournamentID int    `json:"tournament_id"`
	CreatedAt    string `json:"created_at"`
}

type DBMatch struct {
	ID           int    `json:"id"`
	TeamOne      int    `json:"team_one"`
	TeamTwo      int    `json:"team_two"`
	TournamentID int    `json:"tournament_id"`
	Date         string `json:"date"`
	Winner       int    `json:"winner"`
	Score        string `json:"score"`
}

type DBStatsEWC struct {
	MatchID int       `json:"match_id"`
	Stats   GameState `json:"stats"`
}
