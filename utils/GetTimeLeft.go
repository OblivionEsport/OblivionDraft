package utils

import (
	"errors"
	"oblivion/draft/models"
)

func GetTimeLeft(g models.Getter) (int, error) {
	draft, err := GetDraft(g)
	if err != nil {
		return -1, errors.New("getTimeLeft: could not get draft")
	}
	return draft.Timer.AdjustedTimeLeftInPhase, nil
}
