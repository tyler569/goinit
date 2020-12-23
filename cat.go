package main

import (
	"fmt"
	"io/ioutil"
)

func cat(files []string) error {
	for _, file := range files {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		fmt.Print(string(bytes))
	}
	return nil
}
