package controller

import "github.com/gin-gonic/gin"

func CreateUser( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Create User",
	})
}

func GetUser( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get User",
	})
}

func GetUsers( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Get Users",
	})	
}

func UpdateUser( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Update User",
	})	
}

func DeleteUser( c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "Delete User",
	})
}

func signUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Sign Up",
		})
	}
}

func logIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Log In",
		})
	}
}