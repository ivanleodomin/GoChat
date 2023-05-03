package user_handler

import (
	"app-go/internal/platform/storage/postgresql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(userRepository *postgresql.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var userJSON User

		if err := ctx.BindJSON(&userJSON); err != nil {
			fmt.Println(err)
			return
		}

		usr, err := userRepository.Login(userJSON.Email, userJSON.Password)

		if err != nil {
			ctx.String(http.StatusUnauthorized, "Password or mail not found")
			return
		}

		res, err := json.Marshal(usr)

		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error not supported")
			return
		}

		ctx.Data(http.StatusCreated, "application/json", res)
		return
	}
}
