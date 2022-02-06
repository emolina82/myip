// model.go
package model

//Ip data model
type Ip struct {
	Id       string `json:"id"`
	IpAdress string `json:"ip"`
	Country  string `json:"country"`
}

// Ip array
var Ips = []Ip{
	{Id: "1", IpAdress: "192.168.1.200", Country: "Spain"},
	{Id: "2", IpAdress: "192.168.1.220", Country: "Spain"},
}
