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
	if len(token) < 7 {
		models.ErrorResponse(c, http.StatusUnauthorized, "Invalid token", "Token too short")
		return
	}

	claim, err := libs.VerifyJwt(token[7:])
	if err != nil {
		models.ErrorResponse(c, http.StatusUnauthorized, "Invalid token", err.Error())
		return
	}

	var req models.ProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	userId, ok := (*claim)["id"].(float64)
	if !ok {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID in token", "")
		return
	}
	req.UserId = int(userId)

	profile, err := models.EditProfile(c.Request.Context(), req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal update profile", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Sukses mengupdate profile",
		Result:  profile,
	})
}

func UpdateProfileImage(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claim, err := libs.VerifyJwt(token[7:])
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid token", err.Error())
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "gagal ambil gambar", err.Error())
		return
	}

	uid := int((*claim)["id"].(float64))

	// Simpan file ke lokal
	savedFilePath, err := libs.SaveUploadedFile(c, file, "images/user")
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "gagal save file lokal", err.Error())
		return
	}

	// Upload ke cloudinary
	avatarUrl, err := libs.UploadToCloudinary(savedFilePath)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "gagal upload ke cloudinary", err.Error())
		return
	}

	// Update database
	_, err = models.UpdateImage(c, uid, avatarUrl)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "gagal update database", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Update image successfully",
		Result:  avatarUrl,
	})
}

func GetProfileInfo(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	claim, err := libs.VerifyJwt(string(token[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid token", err.Error())
		return
	}

	id := int((*claim)["id"].(float64))

	p, err := models.GetProfileInfo(c, id)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "gagal mendapatkan detail profile", err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.JSON_Response{
		Success: true,
		Message: "Berhasil mendapatkan detail profile",
		Result: gin.H{
			"user_id":  p.UserId,
			"fullname": p.Fullname,
			"image":    p.Image.String,
			"phone":    p.Phone.String,
			"email":    p.Email,
			"address":  p.Address.String,
		},
	})
}
