package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// PhoneHomerAPIv1 configures the Gin Router and sets it to Run
func PhoneHomerAPIv1(url string, port int) bool {
	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		return false
	}
	router.Use(ValidateRequest())
	router.POST("/update", Update)
	router.GET("/getClientList", GetClientList)
	router.POST("/startTunnel", StartTunnel)
	router.POST("/stopTunnel", StopTunnel)
	if err := router.Run(url + ":" + strconv.Itoa(port)); err != nil {
		print(err)
		return false
	}
	return true
}
