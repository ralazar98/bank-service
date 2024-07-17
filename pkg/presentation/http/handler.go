package http

import (
	"bank-service/pkg/infrastructure/memory_cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

type bankServiceI interface {
	// some methods
}

type AccountHandler struct {
	store       memory_cache.BankStorage
	bankService bankServiceI
}

func NewAccountHandler(b memory_cache.BankStorage) *AccountHandler {
	return &AccountHandler{
		store: b,
	}
}

func (a *AccountHandler) Route(g *gin) {
	g.POST("/create", accHandler.CreateAcc)
	g.GET("/list", accHandler.List)
	g.POST("/update", accHandler.UpdateBalance)
	g.POST("/show", accHandler.ShowBalance)
}

func (a *AccountHandler) techRoute(g *gin) {
	//todo: tech
}

func (a *AccountHandler) apiRoute(g *gin) {
	//todo: api
}

func (h *AccountHandler) CreateAcc(c *gin.Context) {
	var account CreateAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.store.Create(account.ID, account.Balance)
}

func (h *AccountHandler) List(c *gin.Context) {
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, r)
}

func (h *AccountHandler) UpdateBalance(c *gin.Context) {
	var account UpdateBalance

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.store.UpdateBalance(account.ID, account.ChangingInBalance, memory_cache.OperationAdd)
	if err != nil {
		if err == memory_cache.NotFoundError {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *AccountHandler) ShowBalance(c *gin.Context) {
	var account ShowBalance
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.store.Show(account.ID)
}
