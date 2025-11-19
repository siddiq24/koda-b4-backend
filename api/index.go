package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/configs"
	"github.com/siddiq24/backend-coffee-shop/routers"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	routers.InitRouter(r)

	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Inisialisasi database
	configs.InitDB()
	configs.InitRedis()

	// Setup router
	router := SetupRouter()

	// Serve request
	router.ServeHTTP(w, r)
}
