package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// repl 交互式命令行
func repl(app *cli.App, appName string) error {
	fmt.Println("input `help` or `h` for help")
	fmt.Println("input `exit` to exit")
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		input.Scan()
		args := strings.Split(input.Text(), " ")
		if isCommandExist(app, args[0]) {
			err := app.Run(append(strings.Split(appName, " "), args...))
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			switch input.Text() {
			case "":
				continue
			case "exit":
				goto BREAK
			case "help", "h":
				_ = app.Run(strings.Split(appName, " -h"))
			default:
				fmt.Println("unrecognized command")
			}
		}
	}
BREAK:
	return nil
}

// isCommandExist 检测命令是否在受支持的列表中
func isCommandExist(app *cli.App, cmd string) bool {
	for _, command := range app.Commands {
		if cmd == command.Name {
			return true
		}
		for _, alias := range command.Aliases {
			if cmd == alias {
				return true
			}
		}
	}
	return false
}
