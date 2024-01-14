package router

import (
	"fmt"
	"net/http"

	"example.com/events-app/models"
	"example.com/events-app/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {

	var user *models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "User created successfully",
	})

}

func loginUser(context *gin.Context) {
	var user *models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	err = user.Login()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "could not authenticate the user",
		})
		return
	}

	token, err := utils.GenerateToken(user.Id, user.Email)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "unable to login user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Logged in successfully",
		"token":   token,
	})

}
