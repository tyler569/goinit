
package main

import (
	"bufio"
    "io/ioutil"
    "fmt"
    "os"
    "os/exec"
    // "strconv"
	"strings"
    "syscall"
)

func errorExit(context string, err error) {
    fmt.Println(context, ":", err)
    for {}
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
    return serial;
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

		str := strings.TrimSpace(s)

		switch {
		case str == "test":
			fmt.Println("Test succesful!")
		case str == "sub":
			sub := exec.Command("/sbin/sub")
			sub.Stdin = os.Stdin
			sub.Stdout = os.Stdout
			sub.Run()
		case len(str) > 3 && str[:3] == "ls ":
			err := listDir(str[3:])
			if err != nil {
				fmt.Println(err)
			}
		case len(str) > 4 && str[:4] == "cat ":
			err := catFile(str[4:])
			if err != nil {
				fmt.Println(err)
			}
		case str == "":
			break
		default:
			fmt.Println("Command not found")
		}
	}
}

