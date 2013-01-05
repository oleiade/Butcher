package main

import (
  "fmt"
  "os"
  "io"
  "bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    file, err := os.Open("/tmp/dat")
    check(err)

    reader := bufio.NewReader(file)

    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF {
            break
        }
        fmt.Println(string(line))
    }
}