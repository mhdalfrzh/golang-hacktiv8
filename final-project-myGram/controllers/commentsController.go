package controllers

import (
	"final-project/database"
	"final-project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	var requestBody struct {
		Message string `json:"message" binding:"required"`
		PhotoID uint   `json:"photo_id"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	Comment.UserID = userID
	Comment.Message = requestBody.Message
	Comment.PhotoID = requestBody.PhotoID

	if err := db.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = userData
	var Comment []models.Comment
	err := db.Debug().Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	commentID, _ := strconv.Atoi(c.Param("commentId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = userData

	var comment models.Comment

	if err := db.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "social media not found",
			"message": err.Error(),
		})
		return
	}

	var updateComment models.Comment
	if err := c.ShouldBindJSON(&updateComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	if updateComment.Message != "" {
		comment.Message = updateComment.Message
	}

	if err := db.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"id":         comment.ID,
		"user_id":    comment.UserID,
		"photo_id":   comment.PhotoID,
		"message":    comment.Message,
		"updated_at": comment.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentID, _ := strconv.Atoi(c.Param("commentId"))
	userData := c.MustGet("userData").(jwt.MapClaims)
	_ = userData

	var comment models.Comment
	if err := db.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "social media not found",
			"message": err.Error(),
		})
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
