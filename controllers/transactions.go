package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/backend-coffee-shop/libs"
	"github.com/siddiq24/backend-coffee-shop/models"
)

type TransactionController struct {
	Transaction models.Transactions
}

func (tc TransactionController) CreateTransactions(c *gin.Context) {
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

func (h *TransactionController) GetHistory(c *gin.Context) {
	claim, err := libs.VerifyJwt((c.Request.Header.Get("Authorization")[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusNetworkAuthenticationRequired, "Unauthorize", err.Error())
		return
	}
	uid := int((*claim)["id"].(float64))

	monthStr := c.DefaultQuery("month", "")
	statusStr := c.DefaultQuery("status", "1")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "5")

	month, err := strconv.Atoi(monthStr)
	status, _ := strconv.Atoi(statusStr)
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}
	if limit <= 5 {
		limit = 5
	}
	if limit >= 10 {
		limit = 10
	}
	if err != nil || monthStr == "" {
		month = int(time.Now().Month())
	}

	req := models.History_req{
		User_id: uid,
		Month:   month,
		Status:  status,
		Page:    page,
		Limit:   limit,
	}

	res, err := h.Transaction.GetHistory(c, req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Service error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Get history successfully",
		Result:  res,
	})
}

func (tc *TransactionController) GetHistoryByInvoice(c *gin.Context) {
	claim, err := libs.VerifyJwt((c.Request.Header.Get("Authorization")[7:]))
	if err != nil {
		models.ErrorResponse(c, http.StatusNetworkAuthenticationRequired, "Unauthorize", err.Error())
		return
	}
	uid := int((*claim)["id"].(float64))
	invoice := c.Param("invoice")
	ress, err := tc.Transaction.GetHistoryByInvoiceID(c.Request.Context(), invoice, uid)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	c.JSON(http.StatusOK, models.JSON_Response{
		Success: true,
		Message: "Get History by invoice successfully",
		Result:  ress,
	})

}
