package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ekonof23/app/database"
	"github.com/ekonof23/app/models"
)

// AuthorizationMiddleware adalah middleware untuk otorisasi pengguna
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dapatkan userID dari context yang disimpan oleh AuthMiddleware
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Dapatkan parameter userID dari URL
		paramUserID := c.Param("userID")

		// Periksa apakah userID di URL sesuai dengan userID pengguna yang terotentikasi
		if userID != paramUserID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
