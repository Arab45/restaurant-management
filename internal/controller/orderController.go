package controller

import "github.com/gin-gonic/gin"

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Create Order",
	})
}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Order",
	})
}
}

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Orders",
	})
}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Update Order",
	})
}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Delete Order",
	})
}
}