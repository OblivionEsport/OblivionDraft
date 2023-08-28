package utils

import (
	"encoding/json"
	"errors"
	"oblivion/draft/models"
	"time"
)

func GetDraft(g models.Getter) (models.Session, error) {
	rawDraft := g.Get("/lol-champ-select/v1/session")
	if rawDraft[2] == 'e' {
		time.Sleep(1 * time.Second)
		return models.Session{}, errors.New("getDraft: Draft not found, try join a game first")
	}
	s := models.Session{}
	err := json.Unmarshal([]byte(rawDraft), &s)
	if err != nil {
		return models.Session{}, err
	}
	return s, nil
}
