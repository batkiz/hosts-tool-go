package main

import (
	"encoding/json"
	"io/ioutil"
)

type hostsItem struct {
	Name    string
	Url     string
	Enabled bool
}

// getHostsItems 由配置文件获得所有配置，并返回为结构体数组
func getHostsItems() []hostsItem {
	var hostsItems []hostsItem
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(content, &hostsItems)

	return hostsItems
}
