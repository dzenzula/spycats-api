package database

import (
	"cmd/catapi/internal/models"
	"errors"

	"gorm.io/gorm"
)

func AddTarget(target *models.AddTarget) error {
	var mission models.Mission

	result := DB.Preload("Targets").First(&mission, target.MissionID)
	if result.Error != nil {
		return result.Error
	}

	if mission.IsComplete {
		return errors.New("cannot add target to a completed mission")
	}

	if len(mission.Targets) == 3 {
		return errors.New("cannot add target, mission has already 3 targets")
	}

	newTarget := models.Target{
		MissionID: mission.ID,
		Name:      target.Name,
		Country:   target.Country,
		Notes:     target.Notes,
	}

	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newTarget).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func RemoveTarget(id int) error {
	var target models.Target
	result := DB.First(&target, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("target not found")
		}
		return result.Error
	}

	if target.IsComplete {
		return errors.New("cannot delete a completed target")
	}

	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&target).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}
