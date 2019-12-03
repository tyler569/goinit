
package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "os/exec"
    "strconv"
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

func listProc() {
    info, err := ioutil.ReadDir("/proc")
    if err != nil {
        errorExit("list proc", err)
    }
    for i, file := range info {
        name := file.Name()
        fmt.Println(i, file.Name())

        if _, err := strconv.Atoi(name); err != nil {
            comm_name := fmt.Sprintf("/proc/%s/comm", name)
            buf := []byte{}
            file, err := os.Open(comm_name)
            if err != nil {
                continue
            }
            file.Read(buf)
            fmt.Println(buf)
        }
    }
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

func serialFile() *os.File {
    serial, err := os.Open("/dev/ttyS0")
    if err != nil {
        errorExit("open serial", err)
    }
    return serial;
}

func main() {
    fmt.Println("Hello World")
    remountRoot()
    mountProc()

    serial := serialFile()

    sub := exec.Command("/sbin/sub")
    sub.Stdin = serial
    sub.Stdout = serial
    sub.Stderr = serial

    sub.Start()
    fmt.Println("running 2")
    sub.Wait()
    fmt.Println("done")

    for {}
}

