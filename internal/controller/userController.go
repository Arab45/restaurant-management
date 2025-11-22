package controller

import "github.com/gin-gonic/gin"



func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get User",
	})
}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Users",
	})	
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update User",
	})	
	}
}


func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Sign Up",
		})
	}
}

func LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Log In",
		})
	}
}