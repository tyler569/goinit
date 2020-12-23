package main

import (
	"fmt"
	"io/ioutil"
)

func ls(dirs []string) error {
	if len(dirs) == 0 {
		dirs = []string{"."}
	}
	for _, dir := range dirs {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, f := range files {
			fmt.Println(f.Name())
		}
	}
	return nil
}
