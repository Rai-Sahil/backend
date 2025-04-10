package transport

import (
	"net/http"

	"github.com/Rai-Sahil/backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Account *service.AccountService
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/open-account", h.OpenAccount)
	router.POST("/deposit", h.Deposit)
	router.POST("/withdraw", h.Withdraw)
	router.GET("/balance/:accountId", h.GetBalance)
}

func (h *Handler) OpenAccount(c *gin.Context) {
	uniqueId := uuid.New().String()
	
	err := h.Account.OpenAccount(uniqueId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetBalance(c *gin.Context) {
	accountId := c.Param("accountId")

	balance, err := h.Account.RebuildBalance(accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (h *Handler) Deposit(c *gin.Context) {
	accountId := c.Query("accountId")
	var body struct {
		Amount int64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Account.Deposit(accountId, body.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Withdraw(c *gin.Context) {
	accountId := c.Query("accountId")
	var body struct {
		Amount int64 `json:"amount"`
	}
	
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Account.Withdraw(accountId, body.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
