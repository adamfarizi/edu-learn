package repository

import (
	"edu-learn/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

type CourseRepository interface {
	CreateCourse(course *model.Course) (model.Course, error)
	GetAllCourse() ([]model.Course, error)
	GetCourseById(id int) (model.Course, error)
	UpdateCourse(id int, course *model.Course) (model.Course, error)
	DeleteCourse(id int) error
}

func (c *courseRepository) CreateCourse(course *model.Course) (model.Course, error) {
	err := c.db.Create(course).Error
	if err != nil {
		return model.Course{}, fmt.Errorf("failed to create course: %w", err)
	}

	return *course, nil
}

func (c *courseRepository) GetAllCourse() ([]model.Course, error) {
	var courses []model.Course

	err := c.db.
		Preload("Instructor").
		Preload("Materials").
		Find(&courses).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get courses: %w", err)
	}

	if len(courses) == 0 {
		return nil, fmt.Errorf("no courses found")
	}

	return courses, nil
}

func (c *courseRepository) GetCourseById(id int) (model.Course, error) {
	var course model.Course

	err := c.db.
		Preload("Instructor").
		Preload("Materials").
		First(&course, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Course{}, fmt.Errorf("course not found")
	}
	if err != nil {
		return model.Course{}, fmt.Errorf("failed to get course: %w", err)
	}

	return course, nil
}

func (c *courseRepository) UpdateCourse(id int, course *model.Course) (model.Course, error) {
	var existingCourse model.Course

	err := c.db.First(&existingCourse, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Course{}, fmt.Errorf("course not found")
	}
	if err != nil {
		return model.Course{}, fmt.Errorf("failed to get course: %w", err)
	}

	err = c.db.
		Model(&existingCourse).
		Updates(course).Error
	if err != nil {
		return model.Course{}, fmt.Errorf("failed to update course: %w", err)
	}

	return existingCourse, nil
}

func (c *courseRepository) DeleteCourse(id int) error {
	var course model.Course

	err := c.db.First(&course, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("course not found")
	}
	if err != nil {
		return fmt.Errorf("failed to get course: %w", err)
	}

	err = c.db.Delete(&course).Error
	if err != nil {
		return fmt.Errorf("failed to delete course: %w", err)
	}

	return nil
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}
