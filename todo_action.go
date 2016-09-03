package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

func todoAction(c *cli.Context) error {
	todoList := os.Getenv("HOME") + "/.paper/todo_list"

	CreateIfNotExistDir(os.Getenv("HOME") + "/.paper")
	CreateIfNotExistFile(todoList)

	if c.NArg() == 0 {
		dat, err := ioutil.ReadFile(todoList)
		Check(err)
		fmt.Print(string(dat))
	} else {
		newContent := c.Args().Get(0) + "\n"
		f, err := os.OpenFile(todoList, os.O_APPEND|os.O_WRONLY, 0666)
		defer f.Close()
		Check(err)
		dat, err := f.WriteString(newContent)
		_ = dat
		Check(err)
	}
	return nil
}
