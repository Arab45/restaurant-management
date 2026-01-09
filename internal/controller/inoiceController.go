package controller

import "github.com/gin-gonic/gin"

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout()
}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Invoice",
	})
}
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Get Invoices",
	})
}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Update Invoice",
	})
}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context){
	c.JSON(200, gin.H{
		"result": "Delete Invoice",
	})
}
}