package controllers

import (
	"assignment-2/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrder(c *gin.Context) {
	var order []models.Order
	models.DB.Find(&order)
	c.JSON(http.StatusOK, gin.H{"order": order})

}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&order).Where("id = ?", id).Updates(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func DeleteOrder(c *gin.Context) {

	var order models.Order

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&order, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
