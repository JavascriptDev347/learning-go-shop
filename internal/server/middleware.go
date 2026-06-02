package server

import (
	"strings"

	"github.com/JavascriptDev347/learning-go-shop/internal/models"
	"github.com/JavascriptDev347/learning-go-shop/internal/utils"
	"github.com/gin-gonic/gin"
)

// authMiddleware func belong to server struct, this function is used to check if the user is authenticated or not, if not authenticated return 401 Unauthorized response, otherwise set user_id, user_email and user_role to context and call next handler
func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthHeader := c.GetHeader("Authorization")
		if AuthHeader == "" {
			utils.UnauthorizedResponse(c, "Authorization header is required")
			c.Abort()
			return
		}

		tokenParts := strings.Split(AuthHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.UnauthorizedResponse(c, "Invalid Authorization header format")
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenParts[1], s.config.JWT.Secret)
		if err != nil {
			utils.UnauthorizedResponse(c, "Invalid token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)
		c.Next()

	}
}

// adminMiddleware for identity admin
func (s *Server) adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			utils.ForbiddenResponse(c, "You are not authorized to access this resource")
			c.Abort()
			return
		}

		if role != string(models.UserRoleAdmin) {
			utils.ForbiddenResponse(c, "You are not authorized to access this resource")
			c.Abort()
			return
		}
		c.Next()
	}
}
