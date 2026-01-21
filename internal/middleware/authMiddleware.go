package middleware

import (
	"github.com/gin-gonic/gin"
	 
)

func AuthMiddleware() gin.HandlerFunc{
	return func (c *gin.Context){
		clientToken := c.Request.Header.get("token")
		if clientToken == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("firstName", claims.First_name)
		c.Set("lastName", claims.Last_name)
		c.Set("uid", claims.Uid)

		c.Next()

	}
}