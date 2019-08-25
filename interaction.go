package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
	"github.com/urfave/cli"
)

func interact(app *cli.App, appName string) error {
	fmt.Println("input `help` or `h` for help")
	fmt.Println("input `exit` to exit")
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">")
		input.Scan()
		if input.Text() == "" {
			continue
		} else if input.Text() == "exit" {
			break
		}
		args := strings.Split(input.Text(), " ")
		if existCommand(app, args[0]) {
			err := app.Run(append(strings.Split(appName, " "), args...))
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			var cmd *exec.Cmd
			if len(args) > 1 {
				cmd = exec.Command(args[0], args[1:]...)
			} else {
				cmd = exec.Command(args[0])
			}
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Println(string(output))
			}
		}
	}
	return nil
}

func existCommand(app *cli.App, cmd string) bool {
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
