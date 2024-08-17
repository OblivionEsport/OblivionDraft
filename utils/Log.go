package utils

import (
	"io"
	"log"
	"os"
	"time"
)

func ConfigLogger() {
	time := time.Now()
	date := time.Format("2006-01-02")
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}
	logFile, err := os.OpenFile("logs/"+date+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logFile.WriteString("\nNew Session\n")
	multi := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multi)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Print("New Logger configured")
}
