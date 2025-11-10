package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/controllers"
	"github.com/siddiq24/backend-coffee-shop/middlewares"
)

func AdminRouter(r *gin.Engine) {
	ac := controllers.AdminController{}
	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware("admin"))
	{
		products := admin.Group("/products")
		products.GET("", ac.GetAllProducts)
		products.POST("", ac.CreateProduct)
		products.PATCH("/:id", ac.UpdateProduct)
		products.DELETE("/:id", ac.DeleteProduct)

		categories := admin.Group("/categories")
		categories.GET("", ac.GetAllCategories)
		categories.POST("", ac.CreateCategory)
		categories.PATCH("/:id", ac.UpdateCategory)
		categories.DELETE("/:id", ac.DeleteCategory)

		transactions := admin.Group("/transactions")
		transactions.GET("", ac.GetAllTransactions)
		transactions.POST("", ac.CreateTransaction)
		transactions.PATCH("/:id", ac.UpdateTransaction)

		users := admin.Group("users")
		users.GET("/users", ac.GetAllUsers)
		users.POST("/users", ac.CreateUser)
		users.PATCH("/users/:id", ac.UpdateUser)
		users.DELETE("/users/:id", ac.DeleteUser)

		images := admin.Group("/products/:id/images")
		images.GET("/:image_id", ac.GetProductImage)
		images.POST("", ac.AddProductImage)
		images.PATCH("/:image_id", ac.UpdateProductImage)
		images.DELETE("/:image_id", ac.DeleteProductImage)
	}
}
