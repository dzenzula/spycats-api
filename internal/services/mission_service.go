package services

import (
	"cmd/catapi/internal/database"
	"cmd/catapi/internal/models"
)

func CreateMission(mission *models.CreateMission) error {
	return database.CreateMission(mission)
}

func RemoveMission(id int) error {
	return database.RemoveMission(id)
}

func UpdateMission(mission *models.UpdateMission) error {
	return database.UpdateMission(mission)
}

func GetMissions() ([]models.Mission, error) {
	return database.GetMissions()
}

func GetMission(id int) (*models.Mission, error) {
	return database.GetMission(id)
}

func AssignCat(missionID int, catID int) error {
	return database.AssignCat(missionID, catID)
}

func UpdateMissionTargets(missionID int, missionTargets []models.UpdateTarget) error {
	return database.UpdateMissionTargets(missionID, missionTargets)
}
