package controller

import (
	"edu-learn/middleware"
	"edu-learn/model"
	"edu-learn/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	useCase        usecase.UserUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (u *userController) Route() {
	// Grup route untuk admin
	adminRoutes := u.rg.Group("/users", u.authMiddleware.RequireToken("admin"))
	{
		adminRoutes.GET("/", u.getAllUsers)
		adminRoutes.DELETE("/:id", u.deleteUser)
	}
	adminUserRoutes := u.rg.Group("/users", u.authMiddleware.RequireToken("admin", "student", "instructor"))
	{
		adminUserRoutes.GET("/:id", u.getUserById)
		adminRoutes.PUT("/:id", u.updateUser)
	}
}

// Method untuk user controller
func (u *userController) getAllUsers(c *gin.Context) {
	user, err := u.useCase.GetAllUsers()
	if err != nil {
		if err.Error() == "no users found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string       `json:"message"`
		Data    []model.User `json:"data"`
	}{
		Message: "User data retrieved successfully",
		Data:    user,
	})
}

func (u *userController) getUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	user, err := u.useCase.GetUserById(userId)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{
		Message: "User data retrieved successfully",
		Data:    user,
	})
}

func (u *userController) updateUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var payload model.User
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.useCase.UpdateUser(userId, &payload)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else if err.Error() == "you may not change your role" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{
		Message: "User updated successfully",
		Data:    user,
	})
}

func (u *userController) deleteUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	err = u.useCase.DeleteUser(userId)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func NewUserController(useCase usecase.UserUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *userController {
	return &userController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
