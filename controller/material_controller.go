package controller

import (
	"edu-learn/middleware"
	"edu-learn/model"
	"edu-learn/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type materialController struct {
	useCase        usecase.MaterialUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (m *materialController) Route() {
	instructorRoutes := m.rg.Group("/courses/:id", m.authMiddleware.RequireToken("instructor"))
	{	
		instructorRoutes.POST("/materials", m.createMaterial)
		instructorRoutes.PUT("/materials/:material_id", m.updateMaterial)
		instructorRoutes.DELETE("materials/:material_id", m.deleteMaterial)
	}
	studentInstructorRoutes := m.rg.Group("/courses/:id", m.authMiddleware.RequireToken("student", "instructor"))
	{
		studentInstructorRoutes.GET("/materials", m.getAllMaterial)
		studentInstructorRoutes.GET("/materials/:material_id", m.getMaterialById)
	}
}

func (m *materialController) getAllMaterial(ctx *gin.Context) {
	courseParam := ctx.Param("id")
	courseId, err := strconv.Atoi(courseParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	materials, err := m.useCase.GetAllMaterial(courseId)
	if err != nil {
		if err.Error() == "course not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		} else if err.Error() == "no materials found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string           `json:"message"`
		Data    []model.Material `json:"data"`
	}{
		Message: "Material data retrieved successfully",
		Data:    materials,
	})
}

func (m *materialController) createMaterial(ctx *gin.Context) {
	courseParam := ctx.Param("id")
	courseId, err := strconv.Atoi(courseParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	var payload model.Material
	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	material, err := m.useCase.CreateMaterial(courseId, &payload)
	if err != nil {
		if err.Error() == "course not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string         `json:"message"`
		Data    model.Material `json:"data"`
	}{
		Message: "Material create successfully",
		Data:    material,
	})
}

func (m *materialController) getMaterialById(ctx *gin.Context) {
	courseParam := ctx.Param("id")
	courseId, err := strconv.Atoi(courseParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	materialParam := ctx.Param("material_id")
	materialId, err := strconv.Atoi(materialParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material id"})
		return
	}

	material, err := m.useCase.GetMaterialById(courseId, materialId)
	if err != nil {
		if err.Error() == "invalid course id" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		} else if err.Error() == "invalid material id" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material id"})
		} else if err.Error() == "material not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string           `json:"message"`
		Data    model.Material `json:"data"`
	}{
		Message: "Material data retrieved successfully",
		Data:    material,
	})
}

func (m *materialController) updateMaterial(ctx *gin.Context) {
	courseParam := ctx.Param("id")
	courseId, err := strconv.Atoi(courseParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	materialParam := ctx.Param("material_id")
	materialId, err := strconv.Atoi(materialParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material id"})
		return
	}

	var payload model.Material
	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	material, err := m.useCase.UpdateMaterial(courseId, materialId, &payload)
	if err != nil {
		if err.Error() == "course not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		} else if err.Error() == "course not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string         `json:"message"`
		Data    model.Material `json:"data"`
	}{
		Message: "Material create successfully",
		Data:    material,
	})
}

func (m *materialController) deleteMaterial(ctx *gin.Context) {
	courseParam := ctx.Param("id")
	courseId, err := strconv.Atoi(courseParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	materialParam := ctx.Param("material_id")
	materialId, err := strconv.Atoi(materialParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material id"})
		return
	}

	err = m.useCase.DeleteMaterial(courseId, materialId)
	if err != nil {
		if err.Error() == "course not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		} else if err.Error() == "material not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}

func NewMaterialController(useCase usecase.MaterialUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *materialController {
	return &materialController{useCase: useCase, rg: rg, authMiddleware: authMiddleware}
}
