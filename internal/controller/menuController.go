package controller

import( 
	"github.com/gin-gonic/gin"
)

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create Menu",
	})
	}
}

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Menus",
	})
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Menu",
	})
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Menu",
	})
	}
}

func DeleteMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Menu",
	})
	}
}