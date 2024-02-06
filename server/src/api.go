package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func PhoneHomerAPIv1(url string, port int) bool {
	router := gin.Default()

	router.Use(validateRequest())
	router.POST("/update", Update)
	router.GET("/getClientList", GetClientList)
	router.POST("/startTunnel", StartTunnel)
	router.POST("/stopTunnel", StopTunnel)

	err := router.Run(url + ":" + strconv.Itoa(port))
	if err != nil {
		print(err)
		return false
	}
	return true
}
