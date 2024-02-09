package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Update takes an UpdateRequest object and hands it to UpdateClient() defined in client.go. Returns a Client Object.
func Update(context *gin.Context) {
	var request UpdateRequest
	if err := context.BindJSON(&request); err != nil {
		print(err)
		return
	}
	client := UpdateClient(request.Id, context.ClientIP(), request.Error)
	context.JSON(http.StatusOK, *client)
}

// GetClientList returns the current contents of the var clients as an array
func GetClientList(context *gin.Context) {
	context.JSON(http.StatusOK, getClientsAsList())
}

// StartTunnel calls AlterTunnel with status set to true
func StartTunnel(context *gin.Context) {
	AlterTunnel(context, true)
}

// StopTunnel calls AlterTunnel with status set to false
func StopTunnel(context *gin.Context) {
	AlterTunnel(context, false)
}

// AlterTunnel calls UpdateClientStatus with the status passed in by the caller
func AlterTunnel(context *gin.Context, status bool) {
	var request StartTunnelRequest
	if err := context.BindJSON(&request); err != nil {
		fmt.Println(err)
		return
	}
	if success := UpdateClientStatus(request.Id, status); !success {
		ReturnError(context, "Client doesn't exist!", http.StatusBadRequest)
	} else {
		context.JSON(http.StatusOK, Response{})
	}
}
