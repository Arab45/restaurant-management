package controller

import "github.com/gin-gonic/gin"

func CreateOrder( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create Order",
	})
	
}

func GetOrder( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Order",
	})
}

func GetOrders( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Orders",
	})
}

func UpdateOrder( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Order",
	})
	
}

func DeleteOrder( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Order",
	})
	
}