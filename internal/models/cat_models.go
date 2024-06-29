package models

type SpyCat struct {
	ID                int     `gorm:"primaryKey;autoIncrement"`
	Name              string  `gorm:"not null"`
	YearsOfExperience int     `gorm:"not null"`
	Breed             string  `gorm:"not null"`
	Salary            float64 `gorm:"not null"`
}

type UpdateSalaryRequest struct {
	ID     int     `json:"id"`
	Salary float64 `json:"salary"`
}

type CreateSpyCat struct {
	Name              string
	YearsOfExperience int
	Breed             string
	Salary            float64
}
