package router

import (
	"net/http"

	"example.com/events-app/models"
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
