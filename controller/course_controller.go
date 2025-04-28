package controller

import (
	"edu-learn/middleware"
	"edu-learn/model"
	"edu-learn/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type courseController struct {
	useCase        usecase.CourseUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *courseController) Route() {
	// Public
	c.rg.GET("/courses", c.getAllCourses)
	c.rg.GET("/courses/:id", c.getCourseById)

	instructorRoutes := c.rg.Group("/courses", c.authMiddleware.RequireToken("instructor"))
	{
		instructorRoutes.POST("/", c.createCourse)
		instructorRoutes.PUT("/:id", c.updateCourse)
		instructorRoutes.DELETE("/:id", c.deleteCourse)
	}
}

func (c *courseController) getAllCourses(ctx *gin.Context) {
	course, err := c.useCase.GetAllCourse()
	if err != nil {
		if err.Error() == "no courses found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No courses found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string         `json:"message"`
		Data    []model.Course `json:"data"`
	}{
		Message: "Course data retrieved successfully",
		Data:    course,
	})
}

func (c *courseController) createCourse(ctx *gin.Context) {
	var payload model.Course

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := c.useCase.CreateCourse(&payload)
	if err != nil {
		if err.Error() == "instructor not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string       `json:"message"`
		Data    model.Course `json:"data"`
	}{
		Message: "Course create successfully",
		Data:    course,
	})
}

func (c *courseController) getCourseById(ctx *gin.Context) {
	id := ctx.Param("id")
	courseId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	course, err := c.useCase.GetCourseById(courseId)
	if err != nil {
		if err.Error() == "invalid course id" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		} else if err.Error() == "course not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string       `json:"message"`
		Data    model.Course `json:"data"`
	}{
		Message: "Course data retrieved successfully",
		Data:    course,
	})
}

func (c *courseController) updateCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	courseId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	var payload model.Course
	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := c.useCase.UpdateCourse(courseId, &payload)
	if err != nil {
		if err.Error() == "course not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		} else if err.Error() == "instructor not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Instructor not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string       `json:"message"`
		Data    model.Course `json:"data"`
	}{
		Message: "Course updated successfully",
		Data:    course,
	})
}

func (c *courseController) deleteCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	courseId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	err = c.useCase.DeleteCourse(courseId)
	if err != nil {
		if err.Error() == "course not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

func NewCourseController(useCase usecase.CourseUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *courseController {
	return &courseController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
