package controller

import "github.com/gin-gonic/gin"

func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Create Note",
	})
}
}

func GetNotes() gin.HandlerFunc {
	return func( c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Notes",
	})
}
}

func GetNote() gin.HandlerFunc {
	return func( c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Note",
	})
}
}

func UpdateNote() gin.HandlerFunc {
	return func (c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Note",
	})
}
}

func DeleteNote() gin.HandlerFunc {
	return func (c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Note",
	})
}
}