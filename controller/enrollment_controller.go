package controller

import (
	"edu-learn/middleware"
	"edu-learn/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type enrollmentController struct {
	useCase        usecase.EnrollmentUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *enrollmentController) Route() {
	c.rg.POST("/courses/:id/enroll", c.authMiddleware.RequireToken("student"), c.enrollCourse)
}

func (h *enrollmentController) enrollCourse(c *gin.Context) {
	courseIDStr := c.Param("id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
		return
	}

	userID, err := middleware.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	err = h.useCase.EnrollCourse(userID, courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully enrolled in course"})
}

func NewEnrollmentController(useCase usecase.EnrollmentUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *enrollmentController {
	return &enrollmentController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
