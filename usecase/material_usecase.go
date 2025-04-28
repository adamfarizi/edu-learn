package usecase

import (
	"edu-learn/model"
	"edu-learn/repository"
	"fmt"
)

type materialUseCase struct {
	repo       repository.MaterialRepository
	repoCourse repository.CourseRepository
}

type MaterialUseCase interface {
	CreateMaterial(idCourse int, material *model.Material) (model.Material, error)
	GetAllMaterial(idCourse int) ([]model.Material, error)
	GetMaterialById(idCourse int, idMaterial int) (model.Material, error)
	UpdateMaterial(idCourse int, idMaterial int, material *model.Material) (model.Material, error)
	DeleteMaterial(idCourse int, idMaterial int) error
}

func (m *materialUseCase) CreateMaterial(idCourse int, material *model.Material) (model.Material, error) {
	_, err := m.repoCourse.GetCourseById(idCourse)
	if err != nil {
		return model.Material{}, fmt.Errorf("course not found")
	}

	material.CourseID = idCourse

	return m.repo.CreateMaterial(material)
}

func (m *materialUseCase) GetAllMaterial(idCourse int) ([]model.Material, error) {
	_, err := m.repoCourse.GetCourseById(idCourse)
	if err != nil {
		return []model.Material{}, fmt.Errorf("course not found")
	}

	return m.repo.GetAllMaterial(idCourse)
}

func (m *materialUseCase) GetMaterialById(idCourse int, idMaterial int) (model.Material, error) {
	if idCourse <= 0 {
		return model.Material{}, fmt.Errorf("invalid course id")
	}

	if idMaterial <= 0 {
		return model.Material{}, fmt.Errorf("invalid material id")
	}

	return m.repo.GetMaterialById(idCourse, idMaterial)
}

func (m *materialUseCase) UpdateMaterial(idCourse int, idMaterial int, material *model.Material) (model.Material, error) {
	_, err := m.repoCourse.GetCourseById(idCourse)
	if err != nil {
		return model.Material{}, fmt.Errorf("course not found")
	}

	_, err = m.GetMaterialById(idCourse, idMaterial)
	if err != nil {
		return model.Material{}, fmt.Errorf("material not found")
	}

	return m.repo.UpdateMaterial(idCourse, idMaterial, material)
}

func (m *materialUseCase) DeleteMaterial(idCourse int, idMaterial int) error {
	_, err := m.GetMaterialById(idCourse, idMaterial)
	if err != nil {
		return fmt.Errorf("material not found")
	}

	return m.repo.DeleteMaterial(idCourse, idMaterial)
}

func NewMaterialUseCase(repo repository.MaterialRepository, repoCourse repository.CourseRepository) MaterialUseCase {
	return &materialUseCase{repo: repo, repoCourse: repoCourse}
}
