package models

type Target struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	MissionID  int    `gorm:"not null"`
	Name       string `gorm:"not null"`
	Country    string `gorm:"not null"`
	Notes      string
	IsComplete bool `gorm:"not null;default:false"`
}

type AddTarget struct {
	MissionID int
	Name      string
	Country   string
	Notes     string
}

type UpdateTarget struct {
	ID         int
	Notes      string
	IsComplete bool
}
