package utils

import (
	"oblivion/draft/models"
)

// GetTeams returns the teams from a file

func GetTeams() ([]models.Team, error) {
	// read the file
	teams, err := models.ReadTeamsFile()
	if err != nil {
		return nil, err
	}
	return teams.Teams, nil
}

func SaveTeams(teams []models.Team) error {
	teamFile, err := models.ReadTeamsFile()
	if err != nil {
		return err
	}
	err = models.SaveTeamsFile(models.TeamFile{Teams: teams, Selected: teamFile.Selected})
	if err != nil {
		return err
	}
	return nil
}

func AddTeam(team *models.Team) (bool, error) {
	teams, err := GetTeams()
	if err != nil {
		return true, err
	}
	// if the team already exists, replace it
	for i, t := range teams {
		if t.Name == team.Name {
			teams[i] = *team
			err = SaveTeams(teams)
			if err != nil {
				return true, err
			}
			return false, nil
		}
	}
	teams = append(teams, *team)
	err = SaveTeams(teams)
	if err != nil {
		return true, err
	}
	return true, nil
}

func DeleteTeam(name string) error {
	teams, err := GetTeams()
	if err != nil {
		return err
	}
	for i, team := range teams {
		if team.Name == name {
			teams = append(teams[:i], teams[i+1:]...)
			break
		}
	}
	err = SaveTeams(teams)
	if err != nil {
		return err
	}
	return nil
}

func GetSelected() ([]string, error) {
	teamsFile, err := models.ReadTeamsFile()
	if err != nil {
		return nil, err
	}
	return teamsFile.Selected, nil
}

func SetSelected(name1, name2 string) error {
	teamsFile, err := models.ReadTeamsFile()
	if err != nil {
		return err
	}
	teamsFile.Selected = []string{name1, name2}
	err = models.SaveTeamsFile(teamsFile)
	if err != nil {
		return err
	}
	return nil
}

func GetMatchID() (string, error) {
	teams, err := models.ReadTeamsFile()
	if err != nil {
		return "", err
	}
	return teams.MatchID, nil
}

func SetMatchID(matchID string) error {
	teams, err := models.ReadTeamsFile()
	if err != nil {
		return err
	}
	teams.MatchID = matchID
	err = models.SaveTeamsFile(teams)
	if err != nil {
		return err
	}
	return nil
}
