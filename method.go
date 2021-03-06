package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/eddieivan01/nic"

	. "github.com/logrusorgru/aurora"
)

// getHosts 从 url 下载 hosts 文件内容
func getHosts(url string) string {
	if url == "LOCAL" {
		content, err := ioutil.ReadFile(getLocalHostsPath())
		if err != nil {
			panic(err)
		}
		return string(content)
	}

	resp, err := nic.Get(url, nil)
	if err != nil {
		panic(err)
	}

	return string(resp.Text)
}

// isPathExist 检测路径是否存在
func isPathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// backupHosts 备份 hosts 文件
func backupHosts() {
	if isPathExist(hostsFilePath) {
		backupedName := getHostsDirPath() + time.Now().Format("2006-01-02-15-04-05")

		err := os.Rename(hostsFilePath, backupedName+".bak")
		if err != nil {
			panic(err)
		}
	}
}

// getHostsContent 从所有 enabled 的 hosts 源获取 hosts
func getHostsContent(hosts []hostsItem) string {
	content := ""
	for i := range hosts {
		if hosts[i].Enabled {
			content += getHosts(hosts[i].Url)
			fmt.Println(hosts[i].Name + " hosts file got")
		} else {
			fmt.Println(hosts[i].Name + " disabled")
		}
	}

	return content
}

// cleanBak 清除备份文件
func cleanBak() {
	files, err := filepath.Glob(getHostsDirPath() + `*.bak`)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		_ = os.Remove(file)
	}
	fmt.Println("All bak files cleaned.")
}

// update 更新 hosts
func update() {
	backupHosts()
	h := getHostsItems()
	content := getHostsContent(h)
	f, _ := os.OpenFile(hostsFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	_, _ = f.WriteString(content)
	fmt.Println(Green("updated successfully"))
	flushDNS()
}

// openHosts 在浏览器中打开 hosts 源网址
func openHosts(name string) {
	if name == "LOCAL" {
		openURL(getLocalHostsPath())
		return
	}

	h := getHostsItems()
	for _, item := range h {
		if item.Name == name {
			openURL(item.Url)
			return
		}
	}
	fmt.Println(name + " not found")
}

// recoverLastBak 恢复最近的备份文件
/*func recoverLastBak() {
	files, err := filepath.Glob(`C:\Windows\System32\drivers\etc\*.bak`)
	if err != nil {
		panic(err)
	}
	if files == nil {
		log.Fatal("sorry but there are no backup files.")
	}
	lastBak := ""
	for _, b := range files {
		if b >= lastBak {
			lastBak = b
		}
	}

	if isPathExist(getHostsPathWithHosts()) {
		err = os.Remove(getHostsPathWithHosts())
	}
	err = os.Rename(lastBak, getHostsPathWithHosts())
	if err != nil {
		log.Fatal("recover backup file failed")
	}
}*/
