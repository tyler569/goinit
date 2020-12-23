package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"
)

var mount = flag.Bool("mount", true, "Mount filesystems")

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

var legacyCommands = map[string]func([]string)error {
	"echo": echo,
	"ls": ls,
	"cd": cd,
	"cat": cat,
}

func main() {
	fmt.Println("Hello World")
	flag.Parse()
	if *mount {
		remountRoot()
		mountProc()
	}

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

		if command == "" {
			continue
		}

		if fn, ok := legacyCommands[command]; ok {
			err := fn(args)
			if err != nil {
				fmt.Printf("%v: %v\n", command, err)
			}
		} else {
			unknown(command)
		}
	}
}
