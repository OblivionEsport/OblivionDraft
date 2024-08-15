package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"oblivion/draft/models"
	"strings"
)

func GetChampionInfo(id string) (ChampionName string, SplashURL string, IconURL string) {
	var championName, splashUrl, iconURL string
	versions := fetch("https://ddragon.leagueoflegends.com/api/versions.json")
	if versions == "" {
		return "", "img/notfound.png", "img/notfound_square.png"
	}
	latest := strings.Split(strings.Split(versions, "\"")[1], "\"")[0]
	rawChampions := fetch(fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%v/data/fr_FR/champion.json", latest))
	if rawChampions == "" {
		return "", "img/notfound.png", "img/notfound_square.png"
	}
	var champions *models.Champions
	err := json.Unmarshal([]byte(rawChampions), &champions)
	if err != nil {
		log.Fatal(err)
	}
	for _, champion := range champions.Data {
		if champion.Key == fmt.Sprintf("%v", id) {
			championName = champion.Name
			splashUrl = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/img/champion/splash/%v_0.jpg", champion.ID)
			iconURL = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%v/img/champion/%v.png", latest, champion.ID)
			return championName, splashUrl, iconURL
		}
	}
	return "", "img/notfound.png", "img/notfound_square.png"
}

func fetch(url string) string {
	// fetch the url and return the body
	req, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return ""
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(body)
}
