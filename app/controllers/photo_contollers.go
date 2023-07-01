package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ekonof23/app/database"
	"github.com/ekonof23/app/models"
)

// CreatePhoto handles creating a new photo
func CreatePhoto(c *gin.Context) {
	db := database.GetDB(c)

	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

// GetPhotos handles retrieving all photos
func GetPhotos(c *gin.Context) {
	db := database.GetDB(c)

	var photos []models.Photo
	if err := db.Preload("User").Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

// UpdatePhoto handles updating a photo's information
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB(c)

	photoID := c.Param("photoId")
	var photo models.Photo
	if err := db.First(&photo, photoID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updatePhoto models.Photo
	if err := c.ShouldBindJSON(&updatePhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photo.Title = updatePhoto.Title
	photo.Caption = updatePhoto.Caption
	photo.PhotoURL = updatePhoto.PhotoURL

	if err := db.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

// DeletePhoto handles deleting a photo
func DeletePhoto(c *gin.Context) {
	db := database.GetDB(c)

	photoID := c.Param("photoId")
	var photo models.Photo
	if err := db.First(&photo, photoID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := db.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted"})
}
