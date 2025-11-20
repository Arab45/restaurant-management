package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrderItem( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create Order Item",
	})

}

func GetOrderItemByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Order Item",
	})
}
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Order Items",
	})	
	}
}

func ItemByOrder( id string) (OrderItem []primitive.M, error error) {
	return OrderItem, error
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