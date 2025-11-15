package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/configs"
	"github.com/siddiq24/backend-coffee-shop/routers"
)

var router *gin.Engine

func init() {
	configs.InitPostgres()
	configs.InitRedis()
	router = routers.InitRouter()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
