package utils

import (
	"fmt"
	"os"
)

func Setup() {
	fmt.Println(`
   ___  _    _ _     _            ___            __ _   
  / _ \| |__| (_)_ _(_)___ _ _   |   \ _ _ __ _ / _| |_ 
 | (_) | '_ \ | \ V / / _ \ ' \  | |) | '_/ _` + "`" + ` |  _|  _|
  \___/|_.__/_|_|\_/|_\___/_||_| |___/|_| \__,_|_|  \__|
														
Welcome to the Oblivion Draft Setup!
This will create the necessary files and directories for the application to run.

The setup will ask you for your Riot API key, witch is required for the endgame overlay to work.
If you don't have one, you can leave it empty and the overlay will still work, but some features will be disabled.

If you want to use the Supabase integration, you will also need to provide your Supabase URL and Key.
If you don't want to use Supabase, you can leave those fields empty as well.

Please make sure you have the necessary permissions to create files and directories in the current directory.`)

	// check if the teams file exists
	if _, err := os.Stat("./teams.json"); os.IsNotExist(err) {
		file, err := os.Create("teams.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString("{\"teams\": [],\"selected\": [], \"matchID\":\"\", \"fearless\": [[],[]]}")
	}

	//download the latest version of the overlay and admin
	UpdateOverlay()

	// check if the .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		// ask the user their API key and write it to the .env file
		println("Please enter your Riot API key (leave empty if none):")
		var api_key string
		_, err := fmt.Scanln(&api_key)
		if err != nil {
			if err.Error() != "unexpected newline" {
				panic(err)
			}
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
			if err.Error() != "unexpected newline" {
				panic(err)
			}
		}
		file.WriteString("\nSUPABASE_URL=" + supabase_url)

		println("Please enter your Supabase Key: (leave empty if you don't have one)")
		var supabase_key string
		_, err = fmt.Scanln(&supabase_key)
		if err != nil {
			if err.Error() != "unexpected newline" {
				panic(err)
			}
		}
		file.WriteString("\nSUPABASE_KEY=" + supabase_key)
		if supabase_url == "" || supabase_key == "" {
			println("Supabase integration will be disabled.")
			file.WriteString("\nUSE_SUPABASE=FALSE")
		} else {
			println("Supabase integration will be enabled.")
			file.WriteString("\nUSE_SUPABASE=TRUE")
		}
		file.WriteString("\nWAIT_LCU=TRUE")
		println("Setup complete! You can now run the application.")
	}
}
