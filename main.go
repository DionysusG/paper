package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "paper"
	app.Usage = "A piece of paper in your terminal"

	//	var renew string

	//	app.Flags = []cli.Flag{
	//		cli.BoolFlag{
	//			Name:        "new, n",
	//			Hidden:      false,
	//			Usage:       "create a new todo list",
	//			destination: &renew,
	//		},
	//	}

	app.Commands = []cli.Command{
		{
			Name:   "todo",
			Usage:  "A todo list",
			Action: todoAction,
		},
	}

	app.Run(os.Args)
}
