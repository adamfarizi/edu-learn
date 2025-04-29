package middleware

import (
	"edu-learn/model"
	"edu-learn/utils/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	jwtService service.JwtService
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization" binding:"required"`
}

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aH authHeader

		err := c.ShouldBindHeader(&aH)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unautorized"})
			return
		}

		token := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", 1)

		tokenClaim, err := a.jwtService.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unautorized"})
			return
		}

		c.Set("user", model.User{ID: tokenClaim.UserId, Role: tokenClaim.Role})

		validRole := false
		for _, role := range roles {
			if role == tokenClaim.Role {
				validRole = true
				break
			}
		}
		if !validRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbiden Resourse"})
			return
		}

		c.Next()
	}
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}

// ExtractUser => Ambil user dari Context
func ExtractUser(c *gin.Context) (model.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return model.User{}, fmt.Errorf("user not found in context")
	}

	userData, ok := user.(model.User)
	if !ok {
		return model.User{}, fmt.Errorf("invalid user data type")
	}

	return userData, nil
}

// ExtractUserID => Ambil userID dari Context
func ExtractUserID(c *gin.Context) (int, error) {
	user, err := ExtractUser(c)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// ExtractUserRole => Ambil role user dari Context
func ExtractUserRole(c *gin.Context) (string, error) {
	user, err := ExtractUser(c)
	if err != nil {
		return "", err
	}
	return user.Role, nil
}

// IsAdmin => Cek apakah user admin
func IsAdmin(c *gin.Context) bool {
	user, err := ExtractUser(c)
	if err != nil {
		return false
	}
	return user.Role == "admin"
}
