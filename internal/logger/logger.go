package logger

import (
	"log"
	"os"
	"time"
)

func GetLogger() *log.Logger {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	today := time.Now().Format("2006-01-02")
	logFileName := "logs/" + today + ".log"

	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could not open log file: %v", err)
	}

	return log.New(logFile, "", log.LstdFlags)
}
