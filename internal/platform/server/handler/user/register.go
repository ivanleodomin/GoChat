package user_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusCreated, "everything is ok in server!")
	}
}
