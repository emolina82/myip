// repository.go
package repository

import (
	"myip/domain/model"
	"myip/secondary/adapter"
	"myip/secondary/dao"
)

// GetMyIpRepo
func GetMyIpRepo() model.Ip {
	return adapter.GetMyIpAdapt(dao.GetMyIpDao())
}
