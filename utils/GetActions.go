package utils

import (
	"errors"
	"oblivion/draft/models"
)

func GetActions(g models.Getter) ([][]models.Action, error) {
	actions, err := GetDraft(g)
	if err != nil {
		return nil, errors.New("could not get actions")
	}
	return actions.Actions, nil
}
