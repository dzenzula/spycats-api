package database

import (
	"cmd/catapi/internal/models"
	"errors"

	"gorm.io/gorm"
)

func CreateMission(mission *models.CreateMission) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var newTargets []*models.Target
		for _, t := range mission.Target {
			newTarget := models.Target{
				Name:    t.Name,
				Country: t.Country,
				Notes:   t.Notes,
			}
			newTargets = append(newTargets, &newTarget)
		}

		newMission := models.Mission{
			IsComplete: mission.IsComplete,
			Targets:    newTargets,
		}

		if err := DB.Create(&newMission).Error; err != nil {
			return err
		}
		return nil
	})
}

func RemoveMission(id int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var mission models.Mission
		if err := tx.First(&mission, id).Error; err != nil {
			return err
		}
		if mission.CatID != nil {
			return errors.New("mission assigned to a cat")
		}

		if err := tx.Where("mission_id = ?", id).Delete(&models.Target{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&models.Mission{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

func UpdateMission(mission *models.UpdateMission) error {
	var updMission models.Mission
	if err := DB.First(&updMission, mission.ID).Error; err != nil {
		return err
	}
	updMission.IsComplete = mission.IsComplete
	return DB.Save(updMission).Error
}

func GetMissions() ([]models.Mission, error) {
	var missions []models.Mission
	if err := DB.Preload("Targets").Find(&missions).Error; err != nil {
		return nil, err
	}

	catMap := make(map[int]*models.SpyCat)
	var missionIDs []int

	for i := range missions {
		if missions[i].CatID != nil {
			missionIDs = append(missionIDs, *missions[i].CatID)
		}
	}

	var spyCats []models.SpyCat
	if err := DB.Find(&spyCats, missionIDs).Error; err != nil {
		return nil, err
	}

	for i := range spyCats {
		catMap[spyCats[i].ID] = &spyCats[i]
	}

	for i := range missions {
		if missions[i].CatID != nil {
			missions[i].Cat = catMap[*missions[i].CatID]
		}
	}

	return missions, nil
}

func GetMission(id int) (*models.Mission, error) {
	var mission models.Mission
	if err := DB.Preload("Targets").First(&mission, id).Error; err != nil {
		return nil, err
	}

	if mission.CatID != nil {
		var spyCat models.SpyCat
		if err := DB.First(&spyCat, *mission.CatID).Error; err != nil {
			return nil, err
		}
		mission.Cat = &spyCat
	}

	return &mission, nil
}

func AssignCat(missionID int, catID int) error {
	var mission models.Mission
	if err := DB.First(&mission, missionID).Error; err != nil {
		return err
	}
	mission.CatID = &catID
	return DB.Save(&mission).Error
}

func UpdateMissionTargets(missionID int, missionTargets []models.UpdateTarget) error {
	var mission models.Mission
	if err := DB.Preload("Targets").First(&mission, missionID).Error; err != nil {
		return err
	}
	if mission.IsComplete {
		return errors.New("cannot update targets for a completed mission")
	}

	for _, updatedTarget := range missionTargets {
		var target models.Target
		if err := DB.First(&target, updatedTarget.ID).Error; err != nil {
			return err
		}
		if target.IsComplete {
			return errors.New("cannot update a completed target")
		}

		target.Notes = updatedTarget.Notes
		target.IsComplete = updatedTarget.IsComplete

		if err := DB.Save(&target).Error; err != nil {
			return err
		}
	}
	return nil
}
