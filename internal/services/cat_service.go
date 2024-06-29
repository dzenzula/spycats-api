package services

import (
	"cmd/catapi/internal/database"
	"cmd/catapi/internal/models"
)

func CreateSpyCat(spyCat *models.CreateSpyCat) error {
	return database.CreateSpyCat(spyCat)
}

func RemoveSpyCat(id int) error {
	return database.RemoveSpyCat(id)
}

func UpdateSpyCat(id int, newSalary float64) error {
	return database.UpdateSpyCat(id, newSalary)
}

func GetSpyCats() ([]models.SpyCat, error) {
	return database.GetSpyCats()
}

func GetSpyCat(id int) (*models.SpyCat, error) {
	return database.GetSpyCat(id)
}
