package user_handler

import (
	"app-go/internal/platform/storage/postgresql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserList struct {
	ID   string `json:"id" validate:"required"`
	Mail string `json:"email" validate:"required"`
}

func GetAll(userRepository *postgresql.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q := ctx.Query("page")
		page, error := strconv.Atoi(q)

		if error != nil {
			ctx.String(http.StatusBadRequest, "Page is not a number")
		}

		users, _ := userRepository.GetAll(page)
		fmt.Println(users)
		res, err := json.Marshal(users)

		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error not supported")
			return
		}

		ctx.Data(http.StatusCreated, "application/json", res)
		return

	}
}
