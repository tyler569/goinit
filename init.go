package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	// "strconv"
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

func printSelf() {
	buf := make([]byte, 32)
	file, err := os.Open("/proc/self/comm")
	if err != nil {
		errorExit("open comm", err)
	}
	file.Read(buf)
	fmt.Println("self:", string(buf))
}

func serialFile(filename string) *os.File {
	serial, err := os.Open(filename)
	if err != nil {
		errorExit("open serial", err)
	}
	return serial
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
		default:
			err := execFilenameAndArgs(args)
			fmt.Println("exec:", err)
		}
	}
}
