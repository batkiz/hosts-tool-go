package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cheynewallace/tabby"
)

func add(name, url string) {
	item := hostsItem{Name: name, Url: url, Enabled: true}

	var h []hostsItem
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(content, &h)
	h = append(h, item)
	newContent, _ := json.Marshal(h)

	f, _ := os.OpenFile(configFilePath, os.O_WRONLY|os.O_TRUNC, 0666)
	defer f.Close()
	fmt.Println(string(newContent))
	_, _ = f.WriteString(string(newContent))

	fmt.Println(name + " successfully added")
}

func del(name string) {
	h := getHostsItems()
	for i, _ := range h {
		if h[i].Name == name {
			h = append(h[:i], h[i+1:]...)
		}
	}

	newContent, _ := json.Marshal(h)
	f, _ := os.OpenFile(configFilePath, os.O_WRONLY|os.O_TRUNC, 0666)
	defer f.Close()
	_, _ = f.WriteString(string(newContent))

	fmt.Println(name + " successfully removed")
}

func list() {
	h := getHostsItems()
	t := tabby.New()
	t.AddHeader("NAME", "ENABLED", "URL")
	for i := range h {
		t.AddLine(h[i].Name, h[i].Enabled, h[i].Url)
	}
	t.Print()
}
