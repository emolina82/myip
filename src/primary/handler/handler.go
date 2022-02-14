// service.go
package handler

import (
	"myip/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIpsHandler
func GetIpsHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetIps())
}

// GetIpByIdHandler
func GetIpByIdHandler(c *gin.Context) {
	id := c.Param("id")

	ip, err := service.GetIpById(id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, ip)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ip not found"})
}

// PostIPHandler
func PostIPHandler(c *gin.Context) {

	var v interface{}

	// newIp.
	if err := c.BindJSON(&v); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unreconized format"})
		return

	}

	if err := service.PostIp(v); err == nil {
		c.IndentedJSON(http.StatusCreated, v)
		return
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unreconized format"})

}

// DeleteIpByIdHandler
func DeleteIpByIdHandler(c *gin.Context) {
	id := c.Param("id")

	err := service.DeleteIpById(id)

	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "delete"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ip not found"})
}
