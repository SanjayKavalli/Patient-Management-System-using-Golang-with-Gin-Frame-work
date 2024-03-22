package middleware

import (
	services "CurdOperation/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// To verify ApiKey
func Apikeyvalidator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//get header
		ApiKey := ctx.GetHeader("ApiKey")

		if len(ApiKey) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Apikey is Not mentioned"})
			ctx.Abort()
			return
		}
		if ApiKey == services.GetKeyValue("ApiKey") {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": " Invalid Apikey"})
			ctx.Abort()
			return
		}

	}
}
