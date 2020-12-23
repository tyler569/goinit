package main

import (
	"fmt"
)

func echo(args []string) error {
	for _, arg := range args {
		fmt.Print(arg, "")
	}
	return nil
}
