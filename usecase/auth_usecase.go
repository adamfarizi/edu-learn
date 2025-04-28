package usecase

import (
	"edu-learn/model"
	"edu-learn/utils/service"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type authenticationUseCase struct {
	userUseCase UserUseCase
	jwtService  service.JwtService
}

type AuthenticationUseCase interface {
	RegisterUseCase(user *model.User) (model.User, error)
	LoginUseCase(email string, password string) (string, error)
}

func (a *authenticationUseCase) RegisterUseCase(user *model.User) (model.User, error) {
	if user.Role == "admin" || user.Role == "teacher" {
		return model.User{}, fmt.Errorf("only student role is allowed for registration")
	}

	return a.userUseCase.CreateUser(user)
}

func (a *authenticationUseCase) LoginUseCase(email string, password string) (string, error) {
	// Dapat diubah jika user login dengan email
	user, err := a.userUseCase.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Membandingkan hashed password dari database dengan password yang diterima
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Jika password tidak cocok
		return "", fmt.Errorf("password incorrect")
	}

	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewAuthenticationUsecase(uc UserUseCase, jwtService service.JwtService) AuthenticationUseCase {
	return &authenticationUseCase{userUseCase: uc, jwtService: jwtService}
}
