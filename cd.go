package main

import (
	"os"
)

func cd(dir string) error {
	return os.Chdir(dir)
}
