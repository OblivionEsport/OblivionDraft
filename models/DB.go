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
