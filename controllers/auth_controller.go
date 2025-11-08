package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/siddiq24/backend-coffee-shop/libs"
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
	request.Password = libs.Create_Hash(request.Password)

	if a.Model.EmailExist(c, request.Email) > 0 {
		models.ErrorResponse(c, http.StatusConflict, "Silahkan login atau gunakan email lain", "Email telah terdaftar")
		return
	}

	if _, err := a.Model.AddUser(c, request); err != nil {
		models.ErrorResponse(c, http.StatusConflict, "Terjadi kesalahan saat menambahkan user", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Register successfully",
		Result:  request,
	})
}

func (a *AuthController) Login(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if libs.Verify_Hash(req.Password, a.Model.PasswordUser(c, req.Email)) {
		c.JSON(http.StatusCreated, models.JSON_Response{
			Success: true,
			Message: "Register successfully",
			Result:  req,
		})
	}

	models.ErrorResponse(c, http.StatusBadRequest, "Password wrong", "")
}
