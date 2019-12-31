package users

import (
	"fmt"
	"github.com/Blackoutta/bookstore-users-api/domain/users"
	"github.com/Blackoutta/bookstore-users-api/services"
	"github.com/Blackoutta/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	fmt.Println(user)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		//TODO: handle json error
		restErr := errors.NewBadRequestError(err.Error())
		ctx.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(user)
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		ctx.JSON(saveErr.Code, saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func GetUser(ctx *gin.Context) {
	userID, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError(fmt.Sprintf("invalid user ID: %v, it should be a number", userID))
		ctx.JSON(err.Code, err)
		return
	}
	if userID <= 0 {
		err := errors.NewBadRequestError("invalid user ID")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		//TODO: handle user creation error
		ctx.JSON(getErr.Code, getErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
}


