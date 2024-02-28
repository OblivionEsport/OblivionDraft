package main

import (
	"log"
	"oblivion/draft/utils"
	"os"
)

func _main() {
	var err error
	log.Print("Updating draft and overlay")
	if os.Getenv("OS") != "Windows_NT" {
		log.Fatal("This is not a windows system, exiting")
	}
	// remove the old draft.exe (check if it exists first)
	if _, err := os.Stat("draft.exe"); os.IsNotExist(err) {
		log.Print("draft.exe does not exist")
	} else if err != nil {
		log.Fatal(err)
	} else {
		err = os.Remove("draft.exe")
		if err != nil {
			log.Fatal(err)
		}
	}

	// copy the new draft.exe
	err = utils.DownloadFile("draft.exe", "https://github.com/Urbskali/oblivionDraft/releases/latest/download/draft.exe")
	if err != nil {
		log.Fatal(err)
	}
	// move overlay/teams_img to a temporary location
	if _, err := os.Stat("overlay/teams_img"); !os.IsNotExist(err) {
		err = os.Rename("overlay/teams_img", "tmp_teams_img")
		if err != nil {
			log.Fatal(err)
		}
	}

	utils.UpdateOverlay()

	// move the teams_img back
	if _, err := os.Stat("tmp_teams_img"); !os.IsNotExist(err) {
		os.Remove("overlay/teams_img")
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename("tmp_teams_img", "overlay/teams_img")
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Print("Draft updated")
}
