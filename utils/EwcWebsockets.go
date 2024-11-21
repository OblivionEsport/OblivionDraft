package utils

import (
	"encoding/json"
	"log"
	"oblivion/draft/models"
	"os"
	"strings"
	"time"

	supa "github.com/supabase-community/supabase-go"

	"github.com/gorilla/websocket"
)

func ConnectAndLogWebSocket(url string, logFileName string) error {
	// Create the log file
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()

	// Initialize the logger
	logger := log.New(logFile, "", log.LstdFlags)

	// Connect to the WebSocket server
	var conn *websocket.Conn
	for {
		conn, _, err = websocket.DefaultDialer.Dial(url, nil)

		if err != nil {
			print("Error connecting to WebSocket server, retrying in 2 seconds\n")
			time.Sleep(2 * time.Second)
		} else {
			break
		}

	}
	defer conn.Close()

	logger.Println("Connected to WebSocket server:", url)

	saveNextGamestate := false

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseUrl == "" || supabaseKey == "" {
		log.Println("Supabase URL or Key is missing, skipping Supabase routes")
	}

	supabase, err := supa.NewClient(supabaseUrl, supabaseKey, &supa.ClientOptions{})

	if err != nil {
		log.Println("Error creating Supabase client:", err)
	}

	// Listen for messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logger.Println("Error reading message:", err)
			break
		}

		messageString := string(message)

		// if message containt statfeed_event or win, save the next gamestate to a file
		if strings.Contains(messageString, "statfeed_event") && strings.Contains(messageString, "\"event_name\":\"Win\"") {
			saveNextGamestate = true
		}
		if saveNextGamestate && strings.Contains(messageString, "gamestate") {
			saveNextGamestate = false
			logger.Println("Saving gamestate to file")
			logger.Println(messageString)

			var gamestate models.GameState
			err := conn.ReadJSON(&gamestate)
			if err != nil {
				logger.Println("Error reading JSON:", err)
			}

			// get matchid from the db and save the gamestate to the db
			r := supabase.Rpc("get_latest_unfinished_match", "exact", nil)
			if r == "" || r == "null" || r == "[]" {
				logger.Println("Error getting match from db:", r)
				continue
			}

			r = strings.Trim(r, "[]")
			var data models.DBMatch
			err = json.Unmarshal([]byte(r), &data)
			if err != nil {
				logger.Println("Error parsing match data:", err)
				logger.Println("Match data:", r)
			} else if len(r) == 0 {
				logger.Println("No match found")
			} else {
				logger.Println("Match found:", data)
				// save the gamestate to the db
				var stats models.DBStatsEWC
				stats.MatchID = data.ID
				stats.Stats = gamestate

				_, _, err := supabase.From("stats_ewc").Insert(stats, false, "", "", "").Execute()
				if err != nil {
					logger.Println("Error saving gamestate to db:", err)
				}
			}

		}
	}

	return nil
}
