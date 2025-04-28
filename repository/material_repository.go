package repository

import (
	"edu-learn/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type materialRepository struct {
	db *gorm.DB
}

type MaterialRepository interface {
	CreateMaterial(material *model.Material) (model.Material, error)
	GetAllMaterial(idCourse int) ([]model.Material, error)
	GetMaterialById(idCourse int, idMaterial int) (model.Material, error)
	UpdateMaterial(idCourse int, idMaterial int, material *model.Material) (model.Material, error)
	DeleteMaterial(idCourse int, idMaterial int) error
}

func (m *materialRepository) CreateMaterial(material *model.Material) (model.Material, error) {
	err := m.db.Create(material).Error
	if err != nil {
		return model.Material{}, fmt.Errorf("failed to create material: %w", err)
	}

	err = m.db.
		Preload("Course").
		First(material, material.ID).Error
	if err != nil {
		return model.Material{}, fmt.Errorf("failed to get data with course: %w", err)
	}

	return *material, nil
}

func (m *materialRepository) GetAllMaterial(idCourse int) ([]model.Material, error) {
	var materials []model.Material

	err := m.db.
		Preload("Course").
		Find(&materials, "course_id = ?", idCourse).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get materials: %w", err)
	}

	if len(materials) == 0 {
		return nil, fmt.Errorf("no materials found")
	}

	return materials, nil
}

func (m *materialRepository) GetMaterialById(idCourse int, idMaterial int) (model.Material, error) {
	var material model.Material

	err := m.db.
		Preload("Course").
		Where("course_id = ?", idCourse).
		First(&material, idMaterial).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Material{}, fmt.Errorf("material not found")
	}
	if err != nil {
		return model.Material{}, fmt.Errorf("failed to get material: %w", err)
	}

	return material, nil
}

func (m *materialRepository) UpdateMaterial(idCourse int, idMaterial int, material *model.Material) (model.Material, error) {
	var existingMaterial model.Material

	err := m.db.
		Where("course_id = ?", idCourse).
		First(&existingMaterial, idMaterial).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Material{}, fmt.Errorf("material not found")
	}
	if err != nil {
		return model.Material{}, fmt.Errorf("failed to get material: %w", err)
	}

	err = m.db.
		Model(&existingMaterial).
		Updates(material).Error
	if err != nil {
		return model.Material{}, fmt.Errorf("failed to update material: %w", err)
	}

	return existingMaterial, nil
}

func (m *materialRepository) DeleteMaterial(idCourse int, idMaterial int) error {
	var material model.Material

	err := m.db.
		Where("course_id = ?", idCourse).
		First(&material, idMaterial).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("material not found")
	}
	if err != nil {
		return fmt.Errorf("failed to get material: %w", err)
	}

	err = m.db.Delete(&material).Error
	if err != nil {
		return fmt.Errorf("failed to delete material: %w", err)
	}

	return nil
}

func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db: db}
}
