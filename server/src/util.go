package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateRequest() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.GetHeader("x-api-key") == conf.ApiKey {
			context.Next()
		} else {
			returnError(context, "Not Authorized!", http.StatusUnauthorized)
		}
	}
}

func returnError(context *gin.Context, error string, status int) {
	context.IndentedJSON(status, Response{
		Error: error,
	})
	_ = context.AbortWithError(status, gin.Error{})
}
