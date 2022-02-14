// service.go
package service

import (
	"encoding/json"
	"errors"
	"myip/domain/model"
)

// GetIps
func GetIps() []model.Ip {
	return model.Ips
}

// GetIpById
func GetIpById(id string) (model.Ip, error) {

	for _, a := range model.Ips {
		if a.Id == id {
			return a, nil
		}
	}
	return model.Ip{}, errors.New("not found")
}

// PostIp
func PostIp(v interface{}) error {

	if res, err := json.Marshal(v); err == nil {

		var newIp model.Ip

		if err := json.Unmarshal(res, &newIp); err == nil {
			// Add the new ip to the slice.
			model.Ips = append(model.Ips, newIp)

			return nil
		}

	}

	return errors.New("unreconized format")
}

// DeleteIpById
func DeleteIpById(id string) error {
	var newIps = []model.Ip{}

	found := false

	for _, a := range model.Ips {
		if a.Id == id {
			found = true
		} else {
			// Add the new ip to the slice
			newIps = append(newIps, a)
		}
	}

	if !found {
		return errors.New("not found")

	}

	model.Ips = newIps
	return nil
}
