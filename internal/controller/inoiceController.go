package controller

import "github.com/gin-gonic/gin"

func CreateInvoice( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create Invoice",
	})
	
}

func GetInvoice(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Invoice",
	})
	
}

func GetInvoices(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Invoices",
	})
	
}

func UpdateInvoice(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update Invoice",
	})
	
}

func DeleteInvoice(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete Invoice",
	})
	
}