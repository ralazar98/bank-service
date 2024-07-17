package http

import (
	"bank-service/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHandler struct {
	store storage.BankStorage
}

func NewAccountHandler(b storage.BankStorage) AccountHandler {
	return AccountHandler{
		store: b,
	}
}

func (h AccountHandler) CreateAcc(c *gin.Context) {
	var account CreateAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.store.Create(account.ID, account.Balance)
}

func (h AccountHandler) List(c *gin.Context) {
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, r)
}

func (h AccountHandler) UpdateBalance(c *gin.Context) {
	var account UpdateBalance

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.store.UpdateBalance(account.ID, account.ChangingInBalance, account.Operation)
	if err != nil {
		if err == storage.NotFoundError {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h AccountHandler) ShowBalance(c *gin.Context) {
	var account ShowBalance
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.store.Show(account.ID)

}
