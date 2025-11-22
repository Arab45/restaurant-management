package controller

import "github.com/gin-gonic/gin"

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Create Table",
	})
}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Table",
	})
}
}

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Tables",
	})
}
}

func UpdateTable( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Table",
	})
}

func DeleteTable() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Delete Table",
	})	
}		
}