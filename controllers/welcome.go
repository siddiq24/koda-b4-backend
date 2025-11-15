package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/configs"
)

func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":     "Coffee Shop API is running!",
		"status":      "success",
		"environment": os.Getenv("VERCEL_ENV"),
		"postgres":    configs.PgMsg,
		"redis":       configs.RdbMsg,
	})
}

func Health(c *gin.Context) {
	health := gin.H{
		"status": "healthy",
		"app":    "running",
	}

	if configs.Pg != nil && configs.Rdb != nil {
		health["database"] = "connected Postgres & Redis "
	} else {
		health["database"] = "not_configured"
	}

	c.JSON(200, health)
}
