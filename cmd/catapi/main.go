package main

import (
	routes "cmd/catapi/api/v1"
	"cmd/catapi/internal/database"
	"cmd/catapi/internal/logger"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.GetLogger()
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path: %v", err)
	}

	dir := filepath.Dir(execPath)
	dir = filepath.Dir(filepath.Dir(dir))
	err = godotenv.Load(filepath.Join(dir, ".env"))
	if err != nil {
		log.Printf("Warning: Error loading .env from project root: %v", err)
		log.Println("Attempting to load .env from current directory")

		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error: Could not load .env file from any location: %v", err)
		}
	}

	database.ConnectDB()
	defer database.CloseDB()

	r, err := routes.SetupRouter()
	if err != nil {
		log.Fatalf("Error setting up routes: " + err.Error())
		return
	}

	r.Run()
}
