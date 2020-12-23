package main

import (
	"errors"
	"os"
)

var lastCd = "/"

func cd(args []string) error {
	var dir string
	if len(args) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		dir, lastCd = lastCd, cwd
	} else if len(args) == 1 {
		dir = args[0]
	} else {
		return errors.New("cd only takes one argument")
	}

	return os.Chdir(dir)
}
