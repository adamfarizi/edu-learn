package dto

type EnrollmentDto struct {
	StudentID int `gorm:"column:student_id;not null"`
	CourseID  int `gorm:"column:course_id;not null"`
}
