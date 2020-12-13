package main

import (
	"fmt"
	"io/ioutil"
)

func cat(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	fmt.Print(string(bytes))
	return nil
}
