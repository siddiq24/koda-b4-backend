package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type AuthController struct {
	Model *models.Auth
}

func NewAuthController(auth *models.Auth) *AuthController {
	return &AuthController{
		Model: auth,
	}
}

func (a *AuthController) Register(c *gin.Context) {
	var request models.AuthRequest
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if request.Role == "" {
		request.Role = "user"
	}

	if a.Model.UserExist(c, request.Email) {
		models.ErrorResponse(c, http.StatusConflict, "Silahkan login atau gunakan email lain", "Email telah terdaftar")
		return
	}

	id, _ := a.Model.AddUser(c, request)
	log.Println(id)
	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Register successfully",
		Result:  request,
	})
}
