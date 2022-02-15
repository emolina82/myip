// adapter.go
package adapter

import (
	"myip/common/IpJsonClient/dataModel"
	"myip/domain/model"
)

// GetMyIpAdapt
func GetMyIpAdapt(mod dataModel.IPJson) model.Ip {

	return model.Ip{
		IpAdress: mod.Query,
		Country:  mod.Country,
	}
}
