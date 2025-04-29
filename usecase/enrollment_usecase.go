package usecase

import (
	"edu-learn/repository"
	"fmt"
)

type enrollmentUseCase struct {
	repo repository.EnrollmentRepository
}

type EnrollmentUseCase interface {
	IsEnrolled(userID, courseID int) (bool, error)
	EnrollCourse(userID, courseID int) error
}

func (e *enrollmentUseCase) IsEnrolled(userID, courseID int) (bool, error) {
    return e.repo.IsEnrolled(userID, courseID)
}

func (e *enrollmentUseCase) EnrollCourse(userID, courseID int) error {
	alreadyEnrolled, err := e.IsEnrolled(userID, courseID)
	if err != nil {
		return err
	}

	if alreadyEnrolled {
		return fmt.Errorf("user already enrolled in course")
	}

	return e.repo.CreateEnrollment(userID, courseID)
}

func NewEnrollmentUseCase(repo repository.EnrollmentRepository) EnrollmentUseCase {
	return &enrollmentUseCase{repo: repo}
}
