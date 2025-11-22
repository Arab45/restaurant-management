package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	var a int
	var b int

	a = 10
	b = 20

	fmt.Println("Creating Food Item")

	c.JSON(200, gin.H{
		"result": a + b,
	})
}

}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Food",
	})
}
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Foods",
	})
}
}

func UpdateFood( ) gin.HandlerFunc{
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Food",
	})
}
}

func DeleteFood( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Food",
	})
}

func round( num float64) int{
	return int(num + 0.5)
}

func toFixed( num float64, precision int) float64{
	return 0.0234
}
