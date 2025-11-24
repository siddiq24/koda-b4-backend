package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/siddiq24/backend-coffee-shop/configs"
	"github.com/siddiq24/backend-coffee-shop/libs"
	"github.com/siddiq24/backend-coffee-shop/models"
)

func AuthMiddleware(Role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.JSON_Response{
				Success: false,
				Message: "Missing Authorization header",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, models.JSON_Response{
				Success: false,
				Message: "Invalid Authorization header format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		key := "jwt:blacklist:" + tokenString
		fmt.Println("key from middleware : ", key)

		blacklisted, err := configs.GetRedis().Get(c, key).Result()
		if err != nil && err != redis.Nil {
			c.JSON(http.StatusInternalServerError, models.JSON_Response{
				Success: false,
				Message: "Internal server error",
			})
			c.Abort()
			return
		}
		fmt.Println("blacklist :", blacklisted)

		if blacklisted != "" {
			c.JSON(http.StatusUnauthorized, models.JSON_Response{
				Success: false,
				Message: "The access token is invalid. You must log in again.",
			})
			c.Abort()
			return
		}
		claims, err := libs.VerifyJwt(tokenString)
		if err != nil {
			fmt.Println("JWT verification error:", err)
			c.JSON(http.StatusUnauthorized, models.JSON_Response{
				Success: false,
				Message: "Invalid or expired token",
			})
			c.Abort()
			return
		}

		userRole, ok := (*claims)["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, models.JSON_Response{
				Success: false,
				Message: "Invalid token payload",
			})
			c.Abort()
			return
		}

		c.Set("user", claims)

		if Role == "all" {
			c.Next()
			return
		}

		if userRole != Role {
			c.JSON(http.StatusForbidden, models.JSON_Response{
				Success: false,
				Message: fmt.Sprintf("Anda bukan %s ", Role),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
