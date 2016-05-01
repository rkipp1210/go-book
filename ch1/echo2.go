// Echo 2 prints out its command line arguments
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", " "
    for _, arg := range os.Args[1:] {
        s += sep + arg
    }
    fmt.Println(s)
}