package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	photoID, _ := strconv.Atoi(c.Param("photoId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = userData

	var photo models.Photo

	if err := db.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	var updatePhoto models.Photo
	if err := c.ShouldBindJSON(&updatePhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	if updatePhoto.Title != "" {
		photo.Title = updatePhoto.Title
	}
	if updatePhoto.Caption != "" {
		photo.Caption = updatePhoto.Caption
	}
	if updatePhoto.PhotoUrl != "" {
		photo.PhotoUrl = updatePhoto.PhotoUrl
	}

	if err := db.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photo := models.Photo{}
	err := db.Debug().Where("user_id = ?", userID).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoID, _ := strconv.Atoi(c.Param("photoId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = userData

	var photo models.Photo
	if err := db.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	if err := db.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "photo deleted successfully",
	})
}