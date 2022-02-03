// main.go
package main

import (
	"fmt"
	"myip/domain/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//router.SetTrustedProxies([]string{"1.1.1.1"})

	router.GET("/ip", service.GetIps)
	router.GET("/ip/:id", service.GetIpById)
	router.POST("/ip", service.PostIp)
	router.DELETE("/ip/:id", service.DeleteIpById)

	router.Run("0.0.0.0:8080")

	fmt.Println("Finishing")
}
