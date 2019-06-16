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

func getHostsItems() []hostsItem {
	var hostsItems []hostsItem
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(content, &hostsItems)

	return hostsItems
}
