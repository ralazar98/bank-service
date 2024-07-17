package main

import (
	"bank-service/pkg/infrastructure/memory_cache"
	http2 "bank-service/pkg/presentation/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//cfg := config.New()

	store := memory_cache.New()

	r := gin.Default()
	accHandler := http2.NewAccountHandler(store)
	r.POST("/create", accHandler.CreateAcc)
	r.GET("/list", accHandler.List)
	r.POST("/update", accHandler.UpdateBalance)
	r.POST("/show", accHandler.ShowBalance)
	http.ListenAndServe(":8080", r)
}
