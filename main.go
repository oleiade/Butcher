package main

import (
  "fmt"
  "flag"
  "os"
  "io"
  "bufio"
)

// /*
// ** Command line arguments
// */
var input = flag.String("input", "", "Input file")
var output = flag.String("output", "", "Output dir")
var sample = flag.Int("sample", 10, "How many lines by file?")

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    flag.Parse()
    file, err := os.Open(*input)
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