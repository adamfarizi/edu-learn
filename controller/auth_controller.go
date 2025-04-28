package controller

import (
	"edu-learn/model"
	"edu-learn/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authUC usecase.AuthenticationUseCase
	rg     *gin.RouterGroup
}

func (a *authController) Route() {
	a.rg.POST("/register", a.registerController)
	a.rg.POST("/login", a.loginController)
}

func (a *authController) registerController(c *gin.Context) {
	var payload model.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authUC.RegisterUseCase(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{
		Message: "Login Success",
		Data:    user,
	})
}

func (a *authController) loginController(c *gin.Context) {
	var payload model.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dapat diubah jika user login dengan email
	token, err := a.authUC.LoginUseCase(payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "Login Success",
		Token:   token,
	})
}

func NewAuthController(authUc usecase.AuthenticationUseCase, rg *gin.RouterGroup) *authController {
	return &authController{authUC: authUc, rg: rg}
}
