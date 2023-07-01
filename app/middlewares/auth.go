package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ekonof23/app/helpers"
	"github.com/ekonof23/app/models"
)

// AuthMiddleware adalah middleware untuk otentikasi pengguna
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Periksa header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Periksa format Authorization header
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		token := authHeaderParts[1]

		// Verifikasi token dan periksa apakah pengguna sudah login atau signup
		claims, err := helpers.VerifyJWTToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Simpan informasi pengguna yang terverifikasi di context
		c.Set("userID", claims.UserID)

		c.Next()
	}
}
