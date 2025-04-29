package usecase

import (
	"edu-learn/model"
	"edu-learn/repository"
	"fmt"
)

type courseUseCase struct {
	repo     repository.CourseRepository
	repoUser repository.UserRepository
}

type CourseUseCase interface {
	CreateCourse(course *model.Course) (model.Course, error)
	GetAllCourse() ([]model.Course, error)
	GetCourseById(id int) (model.Course, error)
	UpdateCourse(id int, course *model.Course) (model.Course, error)
	DeleteCourse(id int) error
	GetCoursesByUserID(userID int) ([]model.Course, error)
}

func (c *courseUseCase) CreateCourse(course *model.Course) (model.Course, error) {
	idInstructor := course.InstructorID

	_, err := c.repoUser.GetUserById(idInstructor)
	if err != nil {
		return model.Course{}, fmt.Errorf("instructor not found")
	}

	return c.repo.CreateCourse(course)
}

func (c *courseUseCase) GetAllCourse() ([]model.Course, error) {
	return c.repo.GetAllCourse()
}

func (c *courseUseCase) GetCourseById(id int) (model.Course, error) {
	if id <= 0 {
		return model.Course{}, fmt.Errorf("invalid course id")
	}

	return c.repo.GetCourseById(id)
}

func (c *courseUseCase) UpdateCourse(id int, course *model.Course) (model.Course, error) {
	_, err := c.repo.GetCourseById(id)
	if err != nil {
		return model.Course{}, fmt.Errorf("course not found")
	}

	idInstructor := course.InstructorID
	_, err = c.repoUser.GetUserById(idInstructor)
	if err != nil {
		return model.Course{}, fmt.Errorf("instructor not found")
	}

	return c.repo.UpdateCourse(id, course)
}

func (c *courseUseCase) DeleteCourse(id int) error {
	_, err := c.repo.GetCourseById(id)
	if err != nil {
		return fmt.Errorf("course not found")
	}

	return c.repo.DeleteCourse(id)
}

func (c *courseUseCase) GetCoursesByUserID(userID int) ([]model.Course, error) {
	user, err := c.repoUser.GetUserById(userID)
	if err != nil {
		return nil, err
	}

	if user.Role != "student" {
		return []model.Course{}, fmt.Errorf("forbidden: only students can access their courses")
	}

	return c.repo.GetCoursesByUserID(userID)
}

func NewCourseUsecase(repo repository.CourseRepository, repoUser repository.UserRepository) CourseUseCase {
	return &courseUseCase{repo: repo, repoUser: repoUser}
}
