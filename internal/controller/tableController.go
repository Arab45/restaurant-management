package controller

import "github.com/gin-gonic/gin"

func CreateTable( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create Table",
	})
}

func GetTable( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Table",
	})
}

func GetTables( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Tables",
	})
	
}

func UpdateTable( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Table",
	})
}

func DeleteTable( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Table",
	})			
}