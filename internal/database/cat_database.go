package database

import (
	"cmd/catapi/internal/models"
	"errors"
)

func CreateSpyCat(spyCat *models.CreateSpyCat) error {
	newSpyCat := models.SpyCat{
		Name:              spyCat.Name,
		YearsOfExperience: spyCat.YearsOfExperience,
		Salary:            spyCat.Salary,
		Breed:             spyCat.Breed,
	}
	return DB.Create(newSpyCat).Error
}

func RemoveSpyCat(id int) error {
	var mission models.Mission
	if err := DB.Where("cat_id = ?", id).First(&mission).Error; err != nil {
		return err
	}

	if !mission.IsComplete {
		return errors.New("mission assigned to the cat is not completed, cannot delete cat")
	}

	return DB.Unscoped().Delete(&models.SpyCat{}, id).Error
}

func UpdateSpyCat(id int, newSalary float64) error {
	var spyCat models.SpyCat
	if err := DB.First(&spyCat, id).Error; err != nil {
		return err
	}

	spyCat.Salary = newSalary
	return DB.Save(&spyCat).Error
}

func GetSpyCats() ([]models.SpyCat, error) {
	var spyCats []models.SpyCat
	if err := DB.Find(&spyCats).Error; err != nil {
		return nil, err
	}
	return spyCats, nil
}

func GetSpyCat(id int) (*models.SpyCat, error) {
	var spyCat models.SpyCat
	if err := DB.First(&spyCat, id).Error; err != nil {
		return nil, err
	}
	return &spyCat, nil
}
