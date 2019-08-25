package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/eddieivan01/nic"

	. "github.com/logrusorgru/aurora"
)

func getHosts(url string) string {
	resp, err := nic.Get(url, nil)
	if err != nil {
		panic(err)
	}

	return string(resp.Text)
}

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func backupHosts() {
	if PathExist(hostsPath) {
		backupedName := getHostsPath() + time.Now().Format("2006-01-02-15-04-05")

		err := os.Rename(hostsPath, backupedName+".bak")
		if err != nil {
			panic(err)
		}
	}
}

func getHostsContent(hosts []hostsItem) string {
	content := ""
	for i := range hosts {
		if hosts[i].Enabled == true {
			content += getHosts(hosts[i].Url)
			fmt.Println(hosts[i].Name + " hosts file got")
		} else {
			fmt.Println(hosts[i].Name + " disabled")
		}
	}

	return content
}

func cleanBak() {
	files, err := filepath.Glob(`C:\Windows\System32\drivers\etc\*.bak`)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		os.Remove(file)
	}
}

func update() {
	backupHosts()
	h := getHostsItems()
	content := getHostsContent(h)
	f, _ := os.OpenFile(hostsPath, os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	_, _ = f.WriteString(content)
	fmt.Println(Green("updated successfully"))
	flushDNS()
}

func openHosts(name string) {
	h := getHostsItems()
	for _, item := range h {
		if item.Name == name {
			openURL(item.Url)
			return
		}
	}
	fmt.Println(name + " source not found")
}
