package user_handler

import (
	user "app-go/internal"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

func Register(userRepository user.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var userJSON User

		if err := ctx.BindJSON(&userJSON); err != nil {
			fmt.Println(err)
			return
		}

		newUser, err := user.NewUser(
			userJSON.Firstname,
			userJSON.Lastname,
			userJSON.Email,
			userJSON.Password,
		)

		if err != nil {
			ctx.String(http.StatusBadRequest, "User not created")
		}

		userRepository.Register(newUser)

		userJson, err := json.Marshal(newUser)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error marshalling user to JSON")
			return
		}

		ctx.Data(http.StatusCreated, "application/json", userJson)
	}
}
