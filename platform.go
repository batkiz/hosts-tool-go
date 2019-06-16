package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getHostsFilePath() string {
	if runtime.GOOS == "windows" {
		return `C:\Windows\System32\drivers\etc\hosts`
	} else {
		return `/etc/hosts`
	}
}

func getHostsPath() string {
	if runtime.GOOS == "windows" {
		return `C:\Windows\System32\drivers\etc\`
	} else {
		return `/etc/`
	}
}

func getConfigFilePath() string {
	execpath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	configFilePath := filepath.Join(filepath.Dir(execpath), "./config.json")
	return configFilePath
}

func flushDNS() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("ipconfig", "/flushdns")
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(output))
	case "linux":
		cmd := exec.Command("/etc/init.d/nscd", "restart")
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(output))
	default:
		fmt.Println("not supported now, please flush DNS yourself")
	}
}
