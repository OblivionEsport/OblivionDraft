package utils

import (
	"fmt"
	"oblivion/draft/models"
	"os"
)

func SaveDraft(data models.Session, i int) {
	// convert data to json
	json, err := data.ToJSON()
	if err != nil {
		panic(err)
	}
	// id data dir exists
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", 0755)
	}
	// save data to a file
	f, err := os.Create(fmt.Sprintf("data/%d.json", i))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(json)
	if err != nil {
		panic(err)
	}
}
