package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/libs"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type OrderController struct {
	Order *models.Order
}

// CreateOrder godoc
// @Summary			Create new order
// @Description		Membuat order
// @Tags			order
// @Accept			x-www-form-urlencoded
// @Produce			json
// @Param			shipping_id			formData int true "dine in atau pick up"
// @Param			payment_method_id	formData int true "metode pembayaran"
// @Param			no_order			formData int true "nomor order"
// @Param			status_id			formData int true "on progres dll"
// @Param			promo_id			formData int true "id promo"
// @Param			id					formData int true "id product"
// @Param			size_id				formData int true "ukuran(reguler, medium, large)"
// @Param			variant_id			formData int true "varian (ice, hot, dll)"
// @Param			qty					formData int true "total item"
// @Param			total				formData int true "total harga"
// @Security		BearerAuth
// @Success			201  {object}  models.JSON_Response
// @Failure			400  {object}  models.JSON_Response
// @Failure			409  {object}  models.JSON_Response
// @Router			/order [post]
func (oc *OrderController) CreateOrder(c *gin.Context) {
	var req []models.ProductOrder
	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)
	claim, err := libs.VerifyJwt(string(token[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid token", err.Error())
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Request invalid", err.Error())
		return
	}

	UserId := int((*claim)["id"].(float64))

	info, err := oc.Order.CreateOrder(c, UserId, req)
	if err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Gagal Membuat Order", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Create Order Successfully",
		Result:  info,
	})
}
