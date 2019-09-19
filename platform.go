package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// getHostsFilePath 返回 hosts 文件路径（带 hosts）
func getHostsFilePath() string {
	if runtime.GOOS == "windows" {
		return `C:\Windows\System32\drivers\etc\hosts`
	} else {
		return `/etc/hosts`
	}
}

// getHostsPath 返回 hosts 文件路径（不带 hosts）
func getHostsPath() string {
	if runtime.GOOS == "windows" {
		return `C:\Windows\System32\drivers\etc\`
	} else {
		return `/etc/`
	}
}

// getConfigFilePath 返回配置文件路径
func getConfigFilePath() string {
	execpath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	configFilePath := filepath.Join(filepath.Dir(execpath), "./config.json")
	return configFilePath
}

// flushDNS 刷新 DNS 缓存
func flushDNS() {
	var err error
	var output []byte
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
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		fmt.Println("unsupported platform, please open it in browser yourself")
	}
	if err != nil {
		log.Fatal(err)
	}
}
