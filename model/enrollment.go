package model

import "time"

type Enrollment struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	StudentID  int       `gorm:"column:student_id;not null"`
	Student    *User     `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE"`
	CourseID   int       `gorm:"column:course_id;not null"`
	Course     *Course   `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	EnrolledAt time.Time `gorm:"column:enrolled_at;autoCreateTime"`
}
