package utils

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mholt/archiver/v3"
)

func UpdateOverlay() {
	err := os.RemoveAll("admin")
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("overlay")
	if err != nil {
		log.Fatal(err)
	}

	err = DownloadFile("admin.zip", "https://github.com/Urbskali/oblivionDraft/releases/latest/download/admin.zip")
	if err != nil {
		log.Fatal(err)
	}
	err = DownloadFile("overlay.zip", "https://github.com/Urbskali/oblivionDraft/releases/latest/download/overlay.zip")
	if err != nil {
		log.Fatal(err)
	}

	err = archiver.Unarchive("admin.zip", "admin")
	if err != nil {
		log.Fatal(err)
	}
	err = archiver.Unarchive("overlay.zip", "overlay")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove("admin.zip")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove("overlay.zip")
	if err != nil {
		log.Fatal(err)
	}

	// Move the files to the correct location
	err = os.Rename("admin", "tmp_admin")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename("overlay", "tmp_overlay")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename("tmp_admin/admin", "admin")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename("tmp_overlay/overlay", "overlay")
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("tmp_admin")
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("tmp_overlay")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Overlay and admin updated")
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
