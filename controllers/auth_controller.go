package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/libs"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type AuthController struct {
	Auth models.Auth
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
	request.Password = libs.Generate_Hash(request.Password)

	if a.Auth.EmailExist(c, request.Email) > 0 {
		models.ErrorResponse(c, http.StatusConflict, "Silahkan login atau gunakan email lain", "Email telah terdaftar")
		return
	}

	if _, err := a.Auth.AddUser(c, request); err != nil {
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

	if req.Email == "" {
		req.Email = c.PostForm("email")
	}
	if req.Password == "" {
		req.Password = c.PostForm("password")
	}

	if req.Email == "" || req.Password == "" {
		models.ErrorResponse(c, http.StatusBadRequest, "Email and password are required", "")
		return
	}

	id, pass, role, err := a.Auth.PasswordIDUser(c, req.Email)
	if err != nil {
		models.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password", err)
		return
	}

	fmt.Println(req.Password, pass)

	if !libs.Verify_Hash(req.Password, pass) {
		models.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password", "failed hasing")
		return
	}

	token, err := libs.GenerateJwt(id, role)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Login successfully",
		Token:   token,
		Result: gin.H{
			"id":    id,
			"email": req.Email,
			"role":  role,
		},
	})
	token = ""
}

func (a *AuthController) ForgotPassword(c *gin.Context) {
	var req models.ForgotPassword
	if err := c.ShouldBind(&req); err != nil || req.Email == "" && req.Origin == "" {
		models.ErrorResponse(c, http.StatusBadRequest, "Bad request", err)
		return
	}
	pin, err := a.Auth.ForgotPassword(c.Request.Context(), req.Email)
	if err != nil && pin != "" {
		models.ErrorResponse(c, http.StatusTooManyRequests, "Too many request", err)
		return
	}

	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	if a.Auth.EmailExist(c, req.Email) == -1 {
		models.ErrorResponse(c, http.StatusBadRequest, "Request invalid ", "email alredy exist")
		return
	}

	link := fmt.Sprintf("%s?pin=%s&email=%s", req.Origin, pin, req.Email)

	err = libs.SendEmail(req.Email, "Reset Password", pin, link)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err)
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "PIN telah dikirim melalui email, pastikan email aktif",
	})

}

func (a *AuthController) ValidatePin(c *gin.Context) {
	var req models.ForgotPassword
	if c.ShouldBind(&req) != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Bad request", nil)
		return
	}

	if a.Auth.EmailExist(c, req.Email) == -1 {
		models.ErrorResponse(c, http.StatusBadRequest, "Request invalid ", nil)
		return
	}

	if !a.Auth.ValidatePIN(c, req.Email, req.Pin) {
		c.JSON(http.StatusBadRequest, models.JSON_Response{
			Success: false,
			Message: "Invalid pin",
		})
		return
	}
	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Your Pin is valid",
	})
}

func (a *AuthController) SetNewPassword(c *gin.Context) {
	var req models.ForgotPassword

	if c.ShouldBind(&req) != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Bad request", c.ShouldBind(&req))
		return
	}

	req.NewPassword = libs.Generate_Hash(req.NewPassword)

	if !a.Auth.ValidatePIN(c, req.Email, req.Pin) {
		c.JSON(http.StatusBadRequest, models.JSON_Response{
			Success: false,
			Message: "Invalid pin",
		})
		return
	}
	fmt.Println(req)

	if a.Auth.UpdatePassword(c, req) != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Update new password successfully",
	})
}

func (a *AuthController) Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)
	claim, err := libs.VerifyJwt(string(token[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid token", err.Error())
		return
	}

	exp := int64((*claim)["exp"].(float64))

	err = a.Auth.Logout(c, token, time.Until(time.Unix(exp, 0)))
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal servis error", nil)
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Logout successfully",
	})

}
