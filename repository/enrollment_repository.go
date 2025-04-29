package repository

import (
	"edu-learn/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type enrollmentRepository struct {
	db *gorm.DB
}

type EnrollmentRepository interface {
	IsEnrolled(userID, courseID int) (bool, error)
	CreateEnrollment(userID, courseID int) error
}

func (e *enrollmentRepository) IsEnrolled(userID, courseID int) (bool, error) {
    var count int64
    err := e.db.Model(&model.Enrollment{}).
        Where("student_id = ? AND course_id = ?", userID, courseID).
        Count(&count).Error
	
	if count == 0 {
		return false, fmt.Errorf("student enrolled yet: %w", err)
	}

    return true, nil
}

func (e *enrollmentRepository) CreateEnrollment(userID, courseID int) error {
	enrollment := model.Enrollment{
		StudentID:  userID,
		CourseID:   courseID,
		EnrolledAt: time.Now(),
	}

	err := e.db.Create(&enrollment).Error
	if err != nil {
		return fmt.Errorf("failed to enroll: %w", err)
	}

	return nil
}

func NewEnrollmentRepository(db *gorm.DB) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}
