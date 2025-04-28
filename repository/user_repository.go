package repository

import (
	"edu-learn/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user *model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(id int, user *model.User) (model.User, error)
	DeleteUser(id int) error
}

func (u *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	// Dapat diubah jika user login dengan email
	err := u.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Method untuk user repository
func (u *userRepository) CreateUser(user *model.User) (model.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return *user, nil
}

func (u *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User

	err := u.db.
		Preload("Courses").
		Preload("Enrollments").
		Preload("Payments").
		Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("no users found")
	}

	return users, nil
}

func (u *userRepository) GetUserById(id int) (model.User, error) {
	var user model.User

	err := u.db.
		Preload("Courses").
		Preload("Enrollments").
		Preload("Payments").
		First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (u *userRepository) UpdateUser(id int, user *model.User) (model.User, error) {
	var existingUser model.User

	// Cari user berdasarkan ID
	err := u.db.First(&existingUser, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	// Update user
	err = u.db.
		Model(&existingUser).
		Updates(user).Error
	if err != nil {
		return model.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	return existingUser, nil
}

func (u *userRepository) DeleteUser(id int) error {
	var user model.User

	// Cari user berdasarkan ID
	err := u.db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("user not found")
	}
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	// Delete user
	err = u.db.Delete(&user).Error
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
