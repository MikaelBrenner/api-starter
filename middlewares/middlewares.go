package middlewares

import (
	"strings"

	"api-starter/models/service"
	"api-starter/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) < 10 {
			c.JSON(401, gin.H{
				"error": "authentication header is missing",
			})
			c.Abort()
			return
		}
		temp := strings.Split(authHeader, "Bearer")
		if len(temp) < 2 {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		tokenString := strings.TrimSpace(temp[1])
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secretKey := utils.EnvVar("TOKEN_KEY")
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			userservice := service.UserService{}
			user, err := userservice.FindByEmail(email)
			if err != nil {
				c.JSON(406, gin.H{
					"error": "user not found",
				})
				c.Abort()
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.JSON(400, gin.H{
				"error": "token is not valid",
			})
			c.Abort()
			return
		}
	}
}
