// Reads and prints itself :P

package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Thanks to https://gobyexample.com/reading-files
    
    dat, err := ioutil.ReadFile("file_IO.go")
    check(err)
    fmt.Print(string(dat))
    
}
