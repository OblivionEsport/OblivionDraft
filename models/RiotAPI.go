package models

type PuuidResponse struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type MatchTimeline struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		EndOfGameResult string `json:"endOfGameResult"`
		FrameInterval   int    `json:"frameInterval"`
		Frames          []struct {
			Timestamp int `json:"timestamp"`
			Events    []struct {
				Type                    string `json:"type"`
				BuildingType            string `json:"buildingType"`
				TowerType               string `json:"towerType"`
				KillerID                int    `json:"killerId"`
				VictimID                int    `json:"victimId"`
				WinningTeam             int    `json:"winningTeam"`
				TeamID                  int    `json:"teamId"`
				AssistingParticipantIds []int  `json:"assistingParticipantIds"`
				MonsterType             string `json:"monsterType"`
			} `json:"events"`
			ParticipantFrames map[string]struct {
				ParticipantID int `json:"participantId"`
				TotalGold     int `json:"totalGold"`
				DamageStats   struct {
					TotalDamageDoneToChamp int `json:"totalDamageDoneToChampions"`
				} `json:"damageStats"`
			} `json:"participantFrames"`
		} `json:"frames"`
	} `json:"info"`
}

type Match struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		GameDuration int `json:"gameDuration"`
		Participants []struct {
			Puuid         string `json:"puuid"`
			ChampionID    int    `json:"championId"`
			TeamID        int    `json:"teamId"`
			ParticipantID int    `json:"participantId"`
		} `json:"participants"`
		Teams []struct {
			TeamID int  `json:"teamId"`
			Win    bool `json:"win"`
			Bans   []struct {
				ChampionID int `json:"championId"`
			} `json:"bans"`
		} `json:"teams"`
	} `json:"info"`
}

type MatchEndGame struct {
	MatchID         string            `json:"matchId"`
	Duration        int               `json:"duration"`
	TeamStats       []TeamStats       `json:"teamStats"`
	IndividualStats []IndividualStats `json:"individualStats"`
}

type TeamStats struct {
	Win         bool  `json:"win"`
	Kills       int   `json:"kills"`
	Deaths      int   `json:"deaths"`
	Assists     int   `json:"assists"`
	Tower       int   `json:"towerKills"`
	Dragon      int   `json:"dragonKills"`
	ElderDragon int   `json:"elderDragonKills"`
	Barons      int   `json:"baronsKills"`
	GoldFrames  []int `json:"goldFrames"`
	Bans        []int `json:"bans"`
}

type IndividualStats struct {
	ParticipantID int    `json:"participantId"`
	Puuid         string `json:"puuid"`
	Stats         Stats  `json:"stats"`
	ChampID       int    `json:"championId"`
}

type Stats struct {
	DamageDealtToChampions int `json:"damageDealtToChampions"`
}
