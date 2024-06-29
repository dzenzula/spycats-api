package services

import (
	"cmd/catapi/internal/database"
	"cmd/catapi/internal/models"
)

func AddTarget(target *models.AddTarget) error {
	return database.AddTarget(target)
}

func RemoveTarget(id int) error {
	return database.RemoveTarget(id)
}
