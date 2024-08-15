package models

import (
	"encoding/json"
	"os"
)

type Team struct {
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Color   string `json:"color"`
	Score   string `json:"score"`
	LogoUrl string `json:"logoUrl"`
}

type TeamFile struct {
	Teams    []Team   `json:"teams"`
	Selected []string `json:"selected"`
	MatchID  string   `json:"matchID"`
}

func ReadTeamsFile() (TeamFile, error) {
	var teams TeamFile
	f, err := os.ReadFile("./teams.json")
	if err != nil {
		return TeamFile{}, err
	}
	err = json.Unmarshal(f, &teams)
	if err != nil {
		return TeamFile{}, err
	}
	return teams, nil
}
func SaveTeamsFile(teamFile TeamFile) error {
	f, err := os.OpenFile("./teams.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	// Marshal the teamFile to JSON

	b, err := json.MarshalIndent(teamFile, "", "  ")
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}
