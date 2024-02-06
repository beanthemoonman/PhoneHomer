package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(context *gin.Context) {
	var request UpdateRequest
	if err := context.BindJSON(&request); err != nil {
		print(err)
		return
	}
	client := UpdateClient(request.Id, context.ClientIP(), request.Error)
	context.IndentedJSON(http.StatusOK, *client)
}

func GetClientList(context *gin.Context) {
	var ret = make([]Client, 0)
	for _, client := range clients {
		ret = append(ret, *client)
	}
	context.IndentedJSON(http.StatusOK, ret)
}

func StartTunnel(context *gin.Context) {
	AlterTunnel(context, true)
}

func StopTunnel(context *gin.Context) {
	AlterTunnel(context, false)
}

func AlterTunnel(context *gin.Context, status bool) {
	var request StartTunnelRequest
	if err := context.BindJSON(&request); err != nil {
		fmt.Println(err)
		return
	}

	if success := UpdateClientStatus(request.Id, status); !success {
		returnError(context, "Client doesn't exist!", http.StatusBadRequest)
	} else {
		context.IndentedJSON(http.StatusOK, Response{})
	}
}
