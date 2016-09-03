package main

import (
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateIfNotExistFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		Check(err)
		f.Close()
	}
}

func CreateIfNotExistDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0777)
		Check(err)
	}
}
