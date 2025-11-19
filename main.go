package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/siddiq24/backend-coffee-shop/configs"
	_ "github.com/siddiq24/backend-coffee-shop/docs"
	"github.com/siddiq24/backend-coffee-shop/routers"
)

// @title           Backend User Management API
// @version         1.0
// @description     API for user management and authentication
// @host      		localhost:8085
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
// @schemes http https
// @tag.name authentication
// @tag.description Authentication endpoints (register, login, logout)
// @tag.name users
// @tag.description User management endpoints (CRUD operations)
// @BasePath  /
func main() {
	godotenv.Load()
	configs.InitDB()
	configs.InitRedis()

	r := gin.Default()
	routers.InitRouter(r)

	port := os.Getenv("PORT")
	r.Run(fmt.Sprintf(":%s", port))
}
