package usecase

import (
	"edu-learn/model"
	"edu-learn/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	repo repository.UserRepository
}

type UserUseCase interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user *model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserById(id int) (model.User, error)
	UpdateUser(id int, user *model.User) (model.User, error)
	DeleteUser(id int) error
}

func (u *userUseCase) GetUserByEmail(email string) (model.User, error) {
	// Dapat diubah jika user login dengan email
	if email == "" {
		return model.User{}, fmt.Errorf("email are required")
	}

	return u.repo.GetUserByEmail(email)
}

// Method untuk user usecase
func (u *userUseCase) CreateUser(user *model.User) (model.User, error) {
	// Hash password sebelum menyimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password")
	}
	user.Password = string(hashedPassword)

	return u.repo.CreateUser(user)
}

func (u *userUseCase) GetAllUsers() ([]model.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUseCase) GetUserById(id int) (model.User, error) {
	// Dapat diubah jika user login dengan email
	if id == 0 {
		return model.User{}, fmt.Errorf("id are required")
	}

	return u.repo.GetUserById(id)
}

func (u *userUseCase) UpdateUser(id int, user *model.User) (model.User, error) {
	existingUser, err := u.repo.GetUserById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user not found")
	}

	if existingUser.Role != user.Role {
		return model.User{}, fmt.Errorf("you may not change your role")
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return model.User{}, fmt.Errorf("failed to hash password")
		}
		user.Password = string(hashedPassword)
	} else {
		user.Password = existingUser.Password
	}

	return u.repo.UpdateUser(id, user)
}

func (u *userUseCase) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
