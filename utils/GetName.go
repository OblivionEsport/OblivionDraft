package utils

import (
	"fmt"
	"oblivion/draft/models"
	"strings"
)

func GetName(g models.Getter, summonerID string) string {
	if summonerID == "0" || summonerID == "" {
		return "Bot"
	}
	raw := g.Get("/lol-summoner/v1/summoners/" + summonerID)
	if strings.Contains(raw, "RPC_ERROR") {
		fmt.Println("Error getting summoner name")
		return "Bot"
	}
	displayName := strings.Split(strings.Split(raw, "\"displayName\":\"")[1], "\"")[0]
	return displayName
}
