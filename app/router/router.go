package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ekonof23/app/controllers"
)

// Initialize initializes the router
func Initialize(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	users := v1.Group("/users")
	{
		users.POST("/register", controllers.RegisterUser)
		users.PUT("/:userId", controllers.UpdateUser)
		users.DELETE("/:userId", controllers.DeleteUser)
	}

	photos := v1.Group("/photos")
	{
		photos.POST("", controllers.CreatePhoto)
		photos.GET("", controllers.GetPhotos)
		photos.PUT("/:photoId", controllers.UpdatePhoto)
		photos.DELETE("/:photoId", controllers.DeletePhoto)
	}
}
