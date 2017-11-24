package cmd

import (
	"encoding/json"
	"log"
)

type (
	zabbixDiscoveryJson struct {
		Data zabbixDiscoveryData `json:"data"`
	}
	zabbixDiscoveryData []zabbixDiscoveryItem
	zabbixDiscoveryItem map[string]string
)

func (d zabbixDiscoveryData) Json() string {
	j := zabbixDiscoveryJson{
		Data: d,
	}
	jsonBytes, err := json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}
