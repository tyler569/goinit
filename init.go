package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

func errorExit(context string, err error) {
	fmt.Println(context, ":", err)
	for {
	}
}

func remountRoot() {
	fmt.Println("Remounting '/' read-write")
	err := syscall.Mount("/dev/sda", "/", "", syscall.MS_REMOUNT, "")
	if err != nil {
		errorExit("mount root", err)
	}
}

func mountProc() {
	fmt.Println("Mouting /proc")

	err := syscall.Mount("", "/proc", "proc", 0, "")
	if err != nil {
		errorExit("mount proc", err)
	}
}

func unknown(cmd string) {
	fmt.Fprintln(os.Stderr, "Unknown command:", cmd)
}

func main() {
	fmt.Println("Hello World")
	remountRoot()
	mountProc()

	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		s, err := input.ReadString('\n')
		if err != nil {
			errorExit("read string", err)
		}
		s = strings.TrimSpace(s)

		args := strings.Split(s, " ")
		command := args[0]
		args = args[1:]

		switch command {
		case "":
			continue
		case "echo":
			for _, arg := range args {
				fmt.Printf("%v ", arg)
			}
			fmt.Println()
		case "cat":
			for _, arg := range args {
				cat(arg)
			}
		case "ls":
			for _, arg := range args {
				ls(arg)
			}
		default:
			unknown(command)
		}
	}
}
