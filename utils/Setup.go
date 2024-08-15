package utils

import (
	"fmt"
	"os"
)

func Setup() {
	// check if the teams file exists
	if _, err := os.Stat("./teams.json"); os.IsNotExist(err) {
		file, err := os.Create("teams.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString("{\"teams\": [],\"selected\": [], \"matchID\":\"\"}")
	}

	//download the latest version of the overlay and admin
	UpdateOverlay()

	// check if the .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		// ask the user their API key and write it to the .env file
		println("Please enter your Riot API key:")
		var api_key string
		_, err := fmt.Scanln(&api_key)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(".env")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString("API_KEY=" + api_key)

		println("Please enter your Supabase URL: (leave empty if you don't have one)")
		var supabase_url string
		_, err = fmt.Scanln(&supabase_url)
		if err != nil {
			panic(err)
		}
		file.WriteString("\nSUPABASE_URL=" + supabase_url)

		println("Please enter your Supabase Key: (leave empty if you don't have one)")
		var supabase_key string
		_, err = fmt.Scanln(&supabase_key)
		if err != nil {
			panic(err)
		}
		file.WriteString("\nSUPABASE_KEY=" + supabase_key)
	}
}
