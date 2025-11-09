package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/libs"
	"github.com/siddiq24/backend-coffee-shop/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		Token, err := libs.VerifyJwt(token)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, models.JSON_Response{
				Success: false,
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}
		fmt.Println(Token)
		c.Next()
	}
}
