package main

import (
	"assignment-2/controllers"
	"assignment-2/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrder)
	r.PUT("/orders/:orderid", controllers.UpdateOrder)
	r.DELETE("/orders/:orderid", controllers.DeleteOrder)

	r.Run()
}
