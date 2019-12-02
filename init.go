
package main

import (
    "io/ioutil"
    "fmt"
    "log"
    "os"
    "os/exec"
    "strconv"
    "syscall"
)

func remountRoot() {
    fmt.Println("Remounting '/' read-write")
    err := syscall.Mount("/dev/sda", "/", "", syscall.MS_REMOUNT, "")
    if err != nil {
        fmt.Println("error:", err)
        for {}
    }
}

func mountProc() {
    fmt.Println("Mouting /proc")
    /*
    err := os.Mkdir("/proc", 0777)
    if err != nil {
        fmt.Println("error:", err)
        for {}
    }
    */
    var err error

    err = syscall.Mount("", "/proc", "proc", 0, "")
    if err != nil {
        fmt.Println("error:", err)
        for {}
    }
}

func listProc() {
    info, err := ioutil.ReadDir("/proc")
    if err != nil {
        log.Fatal(err)
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

func main() {
    fmt.Println("Hello World")
    remountRoot()
    mountProc()

    // listProc()

    /*
    info, err = ioutil.ReadDir("/dev")
    if err != nil {
        log.Fatal(err)
    }
    for i, file := range info {
        fmt.Println(i, file.Name())
    }
    */

    buf := make([]byte, 32)
    file, err := os.Open("/proc/self/comm")
    if err != nil {
        fmt.Println("error:", err)
        for {}
    }
    file.Read(buf)
    fmt.Println("self:", string(buf))

    serial, err := os.Open("/dev/ttyS0")
    if err != nil {
        fmt.Println("error:", err)
        for {}
    }

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

