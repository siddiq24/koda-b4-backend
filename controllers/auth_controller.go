package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

// Register godoc
// @Summary      Register new user
// @Description  Mendaftarkan pengguna baru ke sistem
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        fullname formData string true "Nama lengkap user"
// @Param        email    formData string true "Alamat email yang valid"
// @Param        password formData string true "Password minimal 8 karakter" minlength(8)
// @Param        role     formData string false "Role user (default: user)" Enums(user, admin)
// @Success      201  {object}  models.JSON_Response
// @Failure      400  {object}  models.JSON_Response
// @Failure      409  {object}  models.JSON_Response
// @Router       /auth/register [post]
func (a *AuthController) Register(c *gin.Context) {
	var request models.AuthRequest
	if err := c.ShouldBind(&request); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}
	if request.Email == "" && request.Fullname == "" && request.Password == "" {
		request = models.AuthRequest{
			Fullname: c.PostForm("fullname"),
			Email:    c.PostForm("email"),
			Password: c.PostForm("password"),
			Role:     c.PostForm("role"),
		}
	}

	fmt.Println(request)
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
	request.Password = ""

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Register successfully",
		Result:  request,
	})
}

// Login godoc
// @Summary      Login user
// @Description  Melakukan proses login untuk user yang sudah terdaftar
// @Tags         auth
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        email    formData string true "Alamat email yang terdaftar"
// @Param        password formData string true "Password" minlength(8)
// @Success      200  {object}  models.JSON_Response
// @Failure      400  {object}  models.JSON_Response
// @Router       /auth/login [post]
func (a *AuthController) Login(c *gin.Context) {
	var req models.AuthRequest
	req.Fullname = "-"
	if err := c.ShouldBind(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if req.Email == "" && req.Password == "" {
		req = models.AuthRequest{
			Email:    c.PostForm("email"),
			Password: c.PostForm("password"),
		}
	}

	id, pass, role := a.Model.PasswordIDUser(c, req.Email)
	token, _ := libs.GenerateJwt(id, role)
	if libs.Verify_Hash(req.Password, pass) {
		c.JSON(http.StatusCreated, models.JSON_Response{
			Success: true,
			Message: "Login successfully",
			Token:   token,
			Result:  req,
		})
		return
	}

	models.ErrorResponse(c, http.StatusBadRequest, "Password wrong", "")
}
