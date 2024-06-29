package models

type Mission struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	CatID      *int
	Cat        *SpyCat   `gorm:"-"`
	IsComplete bool      `gorm:"not null;default:false"`
	Targets    []*Target `gorm:"foreignKey:MissionID"`
}

type CreateMission struct {
	IsComplete bool
	Target     []*AddTarget
}

type UpdateMission struct {
	ID         int
	IsComplete bool
}
