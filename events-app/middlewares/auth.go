package middlewares

import (
	"net/http"
	"strings"

	"example.com/events-app/utils"
	"github.com/gin-gonic/gin"
)

func ValidateToken(context *gin.Context) {
	bearer_token := context.Request.Header.Get("Authorization")
	token := strings.Replace(bearer_token, "Bearer ", "", 1)

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Invalid token, not authorized",
		})
		return
	}

	user_claims, err := utils.VerifyUserToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Not authorized",
		})
		return
	}

	context.Set("user_claims", user_claims)

	context.Next()
}
