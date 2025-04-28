package model

import "time"

type Payment struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	StudentID   int       `gorm:"column:student_id;not null"`
	Student     *User     `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE"`
	CourseID    int       `gorm:"column:course_id;not null"`
	Course      *Course   `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
	Amount      float64   `gorm:"type:decimal(10,2);not null"`
	Status      string    `gorm:"type:varchar(50);not null;check:status IN ('pending', 'completed', 'failed')"`
	PaymentDate time.Time `gorm:"column:payment_date;autoCreateTime"`
}
