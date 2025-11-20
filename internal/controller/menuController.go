package controller

import( 
	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create Menu",
	})

}

func GetMenus(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Menus",
	})

}

func GetMenu(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Menu",
	})

}

func UpdateMenu(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Menu",
	})

}

func DeleteMenu(c *gin.Context) {

	c.JSON(200, gin.H{
		"result": "Delete Menu",
	})

}