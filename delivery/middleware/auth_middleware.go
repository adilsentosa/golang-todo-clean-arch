package middleware

import (
	"net/http"
	"time"
	sharedmodel "todo-clean-arch/shared/shared_model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func DeleteTaskMiddleware(jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieToken, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
		}

		getClaim := &sharedmodel.CustomClaims{}
		getToken, err := jwt.ParseWithClaims(cookieToken, getClaim, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !getToken.Valid || getClaim.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		if c.FullPath() == "/tasks/delete" && getClaim.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "This route is only for admin",
			})
			return
		}
		c.Next()
	}
}
