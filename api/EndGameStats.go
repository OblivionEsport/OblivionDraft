package api

import (
	"encoding/json"
	"oblivion/draft/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func EndGameStats(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.ApiGetter)
	matchID := c.Params("id")
	rawTimeline := g.Getf("/lol/match/v5/matches/%s/timeline", matchID)
	rawMatch := g.Getf("/lol/match/v5/matches/%s/", matchID)
	m := models.Match{}
	t := models.MatchTimeline{}

	if rawTimeline == "" || rawMatch == "" {
		return c.Status(404).JSON(fiber.Map{
			"error": "Connection to Riot API failed",
		})
	}

	json.Unmarshal([]byte(rawTimeline), &t)
	json.Unmarshal([]byte(rawMatch), &m)

	if t.Metadata.MatchID == "" || len(m.Info.Teams) == 0 || len(m.Info.Teams[0].Bans) == 0 || len(m.Info.Teams[1].Bans) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Match not found",
		})
	}

	r := models.MatchEndGame{
		MatchID:  t.Metadata.MatchID,
		Duration: m.Info.GameDuration,
		TeamStats: []models.TeamStats{
			{
				Win:         false,
				Kills:       0,
				Deaths:      0,
				Assists:     0,
				Tower:       0,
				Dragon:      0,
				ElderDragon: 0,
				Barons:      0,
				GoldFrames:  []int{},
				Bans: []int{
					m.Info.Teams[0].Bans[0].ChampionID,
					m.Info.Teams[0].Bans[1].ChampionID,
					m.Info.Teams[0].Bans[2].ChampionID,
					m.Info.Teams[0].Bans[3].ChampionID,
					m.Info.Teams[0].Bans[4].ChampionID,
				},
			},
			{
				Win:         false,
				Kills:       0,
				Deaths:      0,
				Assists:     0,
				Tower:       0,
				Dragon:      0,
				ElderDragon: 0,
				Barons:      0,
				GoldFrames:  []int{},
				Bans: []int{
					m.Info.Teams[1].Bans[0].ChampionID,
					m.Info.Teams[1].Bans[1].ChampionID,
					m.Info.Teams[1].Bans[2].ChampionID,
					m.Info.Teams[1].Bans[3].ChampionID,
					m.Info.Teams[1].Bans[4].ChampionID,
				},
			},
		},
		IndividualStats: []models.IndividualStats{},
	}
	// Loop through each frame, get every event and update the stats, and at the last frame, get the end game stats
	for i, frame := range t.Info.Frames {
		r.TeamStats[0].GoldFrames = append(r.TeamStats[0].GoldFrames, 0)
		r.TeamStats[1].GoldFrames = append(r.TeamStats[1].GoldFrames, 0)
		for _, event := range frame.Events {
			if event.Type == "CHAMPION_KILL" {
				//fmt.Printf("KillerID: %d, VictimID: %d\n", event.KillerID, event.VictimID)
				if event.KillerID > 0 && event.KillerID < 6 { // Team 1
					r.TeamStats[0].Kills++
					r.TeamStats[1].Deaths++
					r.TeamStats[0].Assists += len(event.AssistingParticipantIds)
				} else if event.KillerID > 5 && event.KillerID < 11 { // Team 2
					r.TeamStats[1].Assists += len(event.AssistingParticipantIds)
					r.TeamStats[1].Kills++
					r.TeamStats[0].Deaths++
				}
			}
			if event.Type == "BUILDING_KILL" && event.BuildingType == "TOWER_BUILDING" {
				//fmt.Printf("Tower down, KillerID: %d\n", event.KillerID)
				if event.TeamID == 100 {
					r.TeamStats[1].Tower++
				} else if event.TeamID == 200 {
					r.TeamStats[0].Tower++
				}
			}
			if event.Type == "ELITE_MONSTER_KILL" {
				if event.MonsterType == "DRAGON" {
					if event.KillerID > 0 && event.KillerID < 6 { // Team 1
						r.TeamStats[0].Dragon++
					} else {
						r.TeamStats[1].Dragon++
					}
				}
				if event.MonsterType == "ELDER_DRAGON" {
					if event.KillerID > 0 && event.KillerID < 6 { // Team 1
						r.TeamStats[0].ElderDragon++
					} else {
						r.TeamStats[1].ElderDragon++
					}
				}
				if strings.Split(event.MonsterType, "_")[0] == "BARON" {
					if event.KillerID > 0 && event.KillerID < 6 { // Team 1
						r.TeamStats[0].Barons++
					} else {
						r.TeamStats[1].Barons++
					}
				}

			}
			if event.Type == "GAME_END" {
				r.TeamStats[0].Win = event.WinningTeam == 100
				r.TeamStats[1].Win = event.WinningTeam == 200
				//fmt.Println("Game ended")
			}
		}
		for _, participantFrame := range frame.ParticipantFrames {
			if participantFrame.ParticipantID < 6 {
				r.TeamStats[0].GoldFrames[i] += participantFrame.TotalGold
			} else {
				r.TeamStats[1].GoldFrames[i] += participantFrame.TotalGold
			}
		}
		if i == len(t.Info.Frames)-1 {
			for _, participantFrame := range frame.ParticipantFrames {
				r.IndividualStats = append(r.IndividualStats, models.IndividualStats{
					ParticipantID: participantFrame.ParticipantID,
					ChampID:       m.Info.Participants[participantFrame.ParticipantID-1].ChampionID,
					Stats: models.Stats{
						DamageDealtToChampions: participantFrame.DamageStats.TotalDamageDoneToChamp,
					},
				})
			}
		}

	}

	return c.JSON(r)
}
