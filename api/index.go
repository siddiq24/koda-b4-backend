package handler

import (
	"net/http"

	"github.com/siddiq24/backend-coffee-shop/models"
	"github.com/siddiq24/backend-coffee-shop/routers"
)

var router http.Handler

func init() {
	// Initialize connections (sudah ada singleton di models)
	_ = models.Pg
	_ = models.Rdb

	// Initialize router
	r := routers.InitRouter()
	router = r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
