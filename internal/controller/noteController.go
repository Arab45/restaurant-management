package controller

import "github.com/gin-gonic/gin"

func CreateNote( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create Note",
	})
	
}

func GetNotes( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Notes",
	})
	
}

func GetNote( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Note",
	})
	
}

func UpdateNote( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Note",
	})
	
}

func DeleteNote( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Note",
	})
	
}