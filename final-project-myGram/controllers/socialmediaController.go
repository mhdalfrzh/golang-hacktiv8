package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocialMedia)
		fmt.Println("Data dari request JSON:", SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = userData
	var SocialMedia []models.SocialMedia
	err := db.Debug().Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = userData

	var socialMedia models.SocialMedia

	if err := db.First(&socialMedia, socialMediaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "social media not found",
			"message": err.Error(),
		})
		return
	}

	var updateSocialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&updateSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	if updateSocialMedia.Name != "" {
		socialMedia.Name = updateSocialMedia.Name
	}
	if updateSocialMedia.SocialMediaUrl != "" {
		socialMedia.SocialMediaUrl = updateSocialMedia.SocialMediaUrl
	}

	if err := db.Save(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMediaID, _ := strconv.Atoi(c.Param("socialMediaId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = userData

	var socialMedia models.SocialMedia
	if err := db.First(&socialMedia, socialMediaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "social media not found",
			"message": err.Error(),
		})
		return
	}

	if err := db.Delete(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "social media deleted successfully",
	})
}