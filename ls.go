package main

import (
	"fmt"
	"io/ioutil"
)

func ls(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
	return nil
}
