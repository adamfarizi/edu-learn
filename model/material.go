package model

import "time"

type Material struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	CourseID  int       `gorm:"column:course_id;not null"`
	Course    *Course   `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text"`
	FileURL   string    `gorm:"column:file_url;type:varchar(255)"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}
