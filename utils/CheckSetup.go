package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func CheckSetup() {
	// This function is used to check if the setup is correct
	filename := filepath.Base(os.Args[0])

	fmt.Println("Starting up...")
	if _, err := os.Stat("./ui/"); os.IsNotExist(err) {
		log.Printf("UI directory does not exist,running '%s setup' ...", filename)
		Setup()
	}
	// check if the config file exists
	if _, err := os.Stat("./teams.json"); os.IsNotExist(err) {
		log.Print("Teams file does not exist, creating a new one...")
		file, err := os.Create("teams.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString("{\"teams\": [],\"selected\": [], \"matchID\":\"\"}")
		log.Print("Teams file created")
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file, running '%s setup' ...", filename)
		Setup()
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file, try running '%s setup'", filename)
		}
	}
	if _, check := os.LookupEnv("API_KEY"); !check {
		log.Printf("Invalid .env file, missing API_KEY value, running '%s setup' ...", filename)
		Setup()
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file, try running '%s setup'", filename)
		}
	}
}
