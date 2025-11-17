package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateFood(c *gin.Context) {
	var a int
	var b int

	a = 10
	b = 20

	fmt.Println("Creating Food Item")

	c.JSON(200, gin.H{
		"result": a + b,
	})

}

func GetFood( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Food",
	})
}

func GetFoods( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Foods",
	})
}

func UpdateFood( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Food",
	})
}

func DeleteFood( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Food",
	})
}
