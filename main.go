package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/siddiq24/backend-coffee-shop/configs"
	"github.com/siddiq24/backend-coffee-shop/routers"
)

func main() {
	godotenv.Load()

	pg := configs.InitDbConfigs()
	r := routers.InitRouter(pg)

	port := os.Getenv("PORT")
	r.Run(fmt.Sprintf(":%s", port))
}
