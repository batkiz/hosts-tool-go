package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func interact(app *cli.App, appName string) error {
	fmt.Println("input `help` or `h` for help")
	fmt.Println("input `exit` to exit")
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		input.Scan()
		if input.Text() == "" {
			continue
		} else if input.Text() == "exit" {
			break
		}
		args := strings.Split(input.Text(), " ")
		if isCommandExist(app, args[0]) {
			err := app.Run(append(strings.Split(appName, " "), args...))
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("unrecognized command")
		}
	}
	return nil
}

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
