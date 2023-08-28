package utils

import (
	"oblivion/draft/models"
	"strings"
)

func GetName(g models.Getter, summonerID string) string {
	if summonerID == "0" {
		return "Bot"
	}

	raw := g.Get("/lol-summoner/v1/summoners/" + summonerID)
	displayName := strings.Split(strings.Split(raw, "\"displayName\":\"")[1], "\"")[0]
	return displayName
}
