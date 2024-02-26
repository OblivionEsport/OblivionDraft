package utils

import "os"

func Setup() {
	// check if the teams file exists
	if _, err := os.Stat("./teams.json"); os.IsNotExist(err) {
		file, err := os.Create("teams.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString("{\"teams\": [],\"selected\": []}")
	}

	//download the latest version of the overlay and admin
	UpdateOverlay()
}
