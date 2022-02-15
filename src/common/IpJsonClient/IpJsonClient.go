// IpJsonClient.go
package IpJsonClient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"myip/common/IpJsonClient/config"
	"myip/common/IpJsonClient/dataModel"
	"net/http"
)

func GetIPJson() (dataModel.IPJson, error) {

	var ipJson dataModel.IPJson

	client := &http.Client{}
	req, err := http.NewRequest("GET", config.BaseURL, nil)

	if err != nil {
		return ipJson, errors.New("service not available")
	}

	//Headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return ipJson, errors.New("service not available")
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ipJson, errors.New("error getting body")
	}

	json.Unmarshal(bodyBytes, &ipJson)

	return ipJson, nil
}
