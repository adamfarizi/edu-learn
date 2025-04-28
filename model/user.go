package model

import "time"

// Dapat diubah jika user login dengan email
type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Role      string    `gorm:"type:varchar(50);not null;check:role IN ('student', 'instructor', 'admin')"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Courses     []Course     `gorm:"foreignKey:InstructorID;constraint:OnDelete:CASCADE"`
	Enrollments []Enrollment `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE"`
	Payments    []Payment    `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE"`
}
