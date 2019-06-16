package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

var (
	hostsPath      string = getHostsFilePath()
	configFilePath string = getConfigFilePath()
	Version        string = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "hoststool"
	app.Usage = "a command line tool to manage hosts file"
	app.Author = "batkiz"
	app.Version = Version
	// app.Description = "a command line tool to manage hosts file"

	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"u"},
			// Usage:   "update hosts file",
			Action: func(c *cli.Context) error {
				fmt.Println(Red("ATTENTION: admin/root privilage required."))
				update()
				return nil
			},
		},
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "cleanup backup files",
			Action: func(c *cli.Context) error {
				cleanBak()
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "show the config",
			Action: func(c *cli.Context) error {
				list()
				return nil
			},
		},
		{
			Name:  "add",
			Usage: "`add NAME URL` to add a new hosts source",
			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					fmt.Println("use `add NAME URL` to add a new hosts source")
				} else {
					add(c.Args().First(), c.Args().Get(1))
				}

				return nil
			},
		},
		{
			Name:  "del",
			Usage: "`del NAME` to delete a hosts source",
			Action: func(c *cli.Context) error {
				if c.Args().First() == "" {
					fmt.Println("use `del NAME` to delete a hosts source")
				} else {
					del(c.Args().First())
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
