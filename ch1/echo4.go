package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    fmt.Println("command: ", os.Args[0])
    fmt.Println("Arguments:")
    for i, arg := range os.Args[1:] {
        fmt.Println(strconv.Itoa(i) + " " + arg)
    }
}