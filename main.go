package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const httpAddress = ":3030"

func main() {
	fmt.Println("Server running in ", httpAddress)
	srv := gin.New()
	srv.GET("/health", healthHandler)
}

func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Todo bien bb")
}
