
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("This is a subprocess")
    fmt.Println("My pid is", os.Getpid())
}

