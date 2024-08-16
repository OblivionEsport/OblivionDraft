package utils

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mholt/archiver/v3"
)

func UpdateOverlay() {

	// check if the folder exists - Legacy version
	if _, err := os.Stat("admin"); !os.IsNotExist(err) {
		err := os.RemoveAll("admin")
		if err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat("overlay"); !os.IsNotExist(err) {
		err := os.RemoveAll("overlay")
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat("endgame"); !os.IsNotExist(err) {
		err := os.RemoveAll("endgame")
		if err != nil {
			log.Fatal(err)
		}
	}

	// check if the folder exists - New version
	if _, err := os.Stat("ui"); !os.IsNotExist(err) {
		err := os.RemoveAll("ui")
		if err != nil {
			log.Fatal(err)
		}
	}

	err := DownloadFile("ui.zip", "https://github.com/Urbskali/oblivionDraft/releases/latest/download/ui.zip")
	if err != nil {
		log.Fatal(err)
	}

	err = archiver.Unarchive("ui.zip", "tmp_ui")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove("ui.zip")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename("tmp_ui/ui", "ui")
	if err != nil {
		log.Fatal(err)
	}

	err = os.RemoveAll("tmp_ui")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Overlays are updated")
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
