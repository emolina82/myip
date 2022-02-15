// main.go
package main

import (
	"fmt"
	"myip/primary/handler"

	"github.com/gin-gonic/gin"
)

//http://ip-api.com/json

func main() {

	router := gin.Default()

	router.SetTrustedProxies([]string{"1.1.1.1"})

	router.GET("/ip", handler.GetIpsHandler)
	router.GET("/ip/:id", handler.GetIpByIdHandler)
	router.GET("/ip?mine=:mine", handler.GetMyIpHandler)
	router.POST("/ip", handler.PostIPHandler)
	router.DELETE("/ip/:id", handler.DeleteIpByIdHandler)

	//router.Run("0.0.0.0:8080")
	router.RunTLS("0.0.0.0:8080", "../certs/server.rsa.crt", "../certs/server.rsa.key")

	fmt.Println("Finishing")
}
