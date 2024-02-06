package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ValidateRequest looks for the header x-api-key and checks to see if it matches the ApiKey from the config
func ValidateRequest() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.GetHeader("x-api-key") == conf.ApiKey {
			context.Next()
		} else {
			ReturnError(context, "Not Authorized!", http.StatusUnauthorized)
		}
	}
}

// ReturnError stops the handler chain with a Response object with the error filled out and the HTTP status
func ReturnError(context *gin.Context, error string, status int) {
	context.AbortWithStatusJSON(status, Response{
		Error: error,
	})
}
