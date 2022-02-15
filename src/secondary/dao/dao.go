// dao.go
package dao

import (
	"myip/common/IpJsonClient"
	"myip/common/IpJsonClient/dataModel"
)

// GetMyIpDao
func GetMyIpDao() dataModel.IPJson {

	if ipJson, err := IpJsonClient.GetIPJson(); err == nil {
		return ipJson
	}

	return dataModel.IPJson{}

}
