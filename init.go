package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

func listDir(dir string) error {
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range info {
		fmt.Println(file.Name())
	}
	return nil
}

func catFile(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	fmt.Print(string(bytes))
	return nil
}

var PATH []string = []string{
	"/sbin",
	"/bin",
}

func findFileInPath(filename string) (*os.File, error) {
	var f *os.File
	var err error
	for _, dir := range PATH {
		f, err = os.Open(dir + "/" + filename)
		if err == nil {
			break
		}
	}

	return f, err
}

func execFilenameAndArgs(args []string) error {
	f, err := findFileInPath(args[0])

	if err != nil {
		return err
	}

	cmd := exec.Command(f.Name(), args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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

		switch command {
		case "":
			continue
		case "echo":
			for _, arg := range args[1:] {
				fmt.Printf("%v ", arg)
			}
			fmt.Println()
		case "cat":
			for _, arg := range args[1:] {
				catFile(arg)
			}
		default:
			err := execFilenameAndArgs(args)
			if err != nil {
				fmt.Println("exec:", err)
			}
		}
	}
}
