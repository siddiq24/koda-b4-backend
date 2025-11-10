package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/libs"
	"github.com/siddiq24/backend-coffee-shop/models"
)

// UpdateProfile godoc
// @Summary      Update user information
// @Description  Update existing user information by ID via form data
// @Tags         profile
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        fullname   formData  string  true   "User name"        minlength(3)
// @Param        phone     	formData  string  true   "User phone"
// @Param        address  	formData  string  false  "User address"    minlength(6)
// @Param        image		formData  file    false  "User picture"
// @Security	 BearerAuth
// @Success      200       	{object}  models.JSON_Response
// @Failure      400       	{object}  models.JSON_Response
// @Router       /profile [patch]
func UpdateProfile(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	claim, err := libs.VerifyJwt(string(token[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid token", err.Error())
		return
	}
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "gagal ambil gambar", err.Error())
		return
	}
	imege, err := libs.SaveUploadedFile(c, file, header, c.PostForm("fullname"))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "gagal save file", err.Error())
		return
	}
	req := models.Profile{
		UserId:   int((*claim)["id"].(float64)),
		Fullname: c.PostForm("fullname"),
		Image:    imege,
		Phone:    c.PostForm("phone"),
		Address:  c.PostForm("address"),
	}

	profile, err := models.EditProfile(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "gagal update profile", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Sukses mengupdate profile",
		Result:  profile,
	})
}
