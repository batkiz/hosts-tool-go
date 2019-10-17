package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// getHostsPathWithoutHosts 返回 hosts 文件路径（不带 hosts）
func getHostsPathWithoutHosts() string {
	switch runtime.GOOS {
	case "windows":
		return `C:\Windows\System32\drivers\etc\`
	case "linux":
		return `/etc/`
	default:
		log.Fatal("sorry, this is an unsupported platform.")
		return ""
	}
}

// getHostsPathWithHosts 返回 hosts 文件路径（带 hosts）
func getHostsPathWithHosts() string {
	return getHostsPathWithoutHosts() + "hosts"
}

// getConfigFilePath 返回配置文件路径
func getConfigFilePath() string {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	configFilePath := filepath.Join(filepath.Dir(execPath), "./config.json")
	return configFilePath
}

// getLocalHostsPath 返回本地 hosts 文件的路径
func getLocalHostsPath() string {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	localHostsFilePath := filepath.Join(filepath.Dir(execPath), "./LOCAL.txt")
	if !isPathExist(localHostsFilePath) {
		f, _ := os.OpenFile(localHostsFilePath, os.O_CREATE|os.O_RDONLY, 0666)
		defer f.Close()
	}
	return localHostsFilePath
}

// flushDNS 刷新 DNS 缓存
func flushDNS() {
	var (
		err    error
		output []byte
	)
	switch runtime.GOOS {
	case "windows":
		output, err = exec.Command("ipconfig", "/flushdns").Output()
		fmt.Println(string(output))
	case "linux":
		output, err = exec.Command("/etc/init.d/nscd", "restart").Output()
		fmt.Println(string(output))
	default:
		fmt.Println("not supported now, please flush DNS yourself")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// openURL 在浏览器中打开 url
func openURL(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		fmt.Println("unsupported platform, please open it manually\n" + url)
	}
	if err != nil {
		log.Fatal(err)
	}
}
