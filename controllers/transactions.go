package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/libs"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type TransactionControoller struct {
	Transaction models.Transactions
}

func (tc TransactionControoller) CreateTransactions(c *gin.Context) {
	claim, err := libs.VerifyJwt((c.Request.Header.Get("Authorization")[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusNetworkAuthenticationRequired, "Unauthorize", err.Error())
		return
	}
	id := int((*claim)["id"].(float64))

	var req models.Transaction_Request
	if err := c.ShouldBind(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}
	req.UserId = id

	if err := tc.Transaction.CreateTransactions(c.Request.Context(), req); err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Server error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "create transactions successfull",
	})
}
