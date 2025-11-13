package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/models"
	"github.com/siddiq24/backend-coffee-shop/routers"
)

var router *gin.Engine

func init() {
	router = routers.InitRouter()
	router.GET("/", func(c *gin.Context) {
		dbStatus := "disconnected"
		rdbStatus := "disconnected"
		if models.GetDB() != nil {
			dbStatus = "connected"
		}
		if models.GetRedis() != nil {
			rdbStatus = "connected"
		}

		c.JSON(200, gin.H{
			"message":     "Coffee Shop API is running!",
			"status":      "success",
			"environment": os.Getenv("VERCEL_ENV"),
			"postgres":    dbStatus,
			"redis":       rdbStatus,
		})
	})

	router.GET("/health", func(c *gin.Context) {
		health := gin.H{
			"status": "healthy",
			"app":    "running",
		}

		if db := models.GetDB(); db != nil {
			health["database"] = "connected"
		} else {
			health["database"] = "not_configured"
		}

		c.JSON(200, health)
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
