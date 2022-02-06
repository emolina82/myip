// service.go
package service

import (
	"myip/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIps
func GetIps(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.Ips)
}

// GetIpById
func GetIpById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range model.Ips {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ip not found"})
}

// PostIp
func PostIp(c *gin.Context) {
	var newIp model.Ip

	// newIp.
	if err := c.BindJSON(&newIp); err != nil {
		return
	}

	// Add the new ip to the slice.
	model.Ips = append(model.Ips, newIp)
	c.IndentedJSON(http.StatusCreated, newIp)
}

// DeleteIpById
func DeleteIpById(c *gin.Context) {
	var newIps = []model.Ip{}

	found := false

	id := c.Param("id")

	for _, a := range model.Ips {
		if a.Id == id {
			found = true
		} else {
			// Add the new ip to the slice
			newIps = append(newIps, a)
		}
	}

	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ip not found"})
		return

	}

	model.Ips = newIps
	c.IndentedJSON(http.StatusOK, gin.H{"message": "delete"})
}
