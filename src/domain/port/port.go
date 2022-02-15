// port.go
package port

import (
	"myip/domain/model"
	"myip/secondary/repository"
)

// GetMyIpPort
func GetMyIpPort() model.Ip {
	return repository.GetMyIpRepo()
}
