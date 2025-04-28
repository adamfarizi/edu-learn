package model

import "time"

type Course struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Title        string    `gorm:"type:varchar(255);not null;unique"`
	Description  string    `gorm:"type:text"`
	InstructorID int       `gorm:"column:instructor_id;not null" json:"instructor_id"` // jika field lebih dari 1 kata maka ditambahi json
	Instructor   *User     `gorm:"foreignKey:InstructorID;constraint:OnDelete:CASCADE"`
	Price        float64   `gorm:"type:decimal(10,2);default:0"`
	Category     string    `gorm:"type:varchar(100)"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Materials []Material `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
}
