package controller

import "github.com/gin-gonic/gin"

func CreateOrderItem( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create Order Item",
	})

}

func GetOrderItem( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Order Item",
	})
	
}

func GetOrderItems( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Order Items",
	})	
}

func UpdateOrderItem( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Order Item",
	})
	
}

func DeleteOrderItem( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Order Item",
	})
	
}