// Echo 3 prints out its command line arguments
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("command: ", os.Args[0])
    fmt.Println("arguments: ", strings.Join(os.Args[1:], " "))
}