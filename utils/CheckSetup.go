package utils

import (
	"fmt"
	"log"
	"os"
)

func CheckSetup() {
	// This function is used to check if the setup is correct

	fmt.Println("Starting up...")
	// check if the files directory exists
	if _, err := os.Stat("./overlay"); os.IsNotExist(err) {
		log.Fatal("Overlay directory does not exist, try running 'draft setup'")
	}
	// check if the ui directory exists
	if _, err := os.Stat("./admin"); os.IsNotExist(err) {
		log.Fatal("Admin directory does not exist, try running 'draft setup'")
	}
	// check if the config file exists
	if _, err := os.Stat("./teams.json"); os.IsNotExist(err) {
		log.Print("Teams file does not exist, creating a new one...")
		file, err := os.Create("teams.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString("{\"teams\": [],\"selected\": []}")
		log.Print("Teams file created")
	}

}
