// 对 config.json 文件的操作
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cheynewallace/tabby"
)

// add 添加源
func add(name, url string) {
	item := hostsItem{Name: name, Url: url, Enabled: true}

	h := getHostsItems()
	h = append(h, item)
	newContent, _ := json.Marshal(h)

	f, _ := os.OpenFile(configFilePath, os.O_WRONLY|os.O_TRUNC, 0666)
	defer f.Close()
	fmt.Println(string(newContent))
	_, _ = f.WriteString(string(newContent))

	fmt.Println(name + " successfully added")
}

// del 删除源
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

// list 输出配置文件信息
func list() {
	h := getHostsItems()
	t := tabby.New()
	t.AddHeader("NAME", "ENABLED", "URL")
	for i := range h {
		t.AddLine(h[i].Name, h[i].Enabled, h[i].Url)
	}
	t.Print()
}
