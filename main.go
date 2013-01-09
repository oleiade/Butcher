package main

import (
  "log"
  "flag"
  "os"
  "io"
  "bufio"
)

// /*
// ** Command line arguments
// */
var input string
var output string
var sample string

func init() {
  const (
    defaultInput = "Input file"
    defaultOutput = "Ouput directory"
    defaultSample = "Line count per output file"
    usage = "butcher -input bigfatfile -output dir -sample linecountpersubfile"
  )

  flag.StringVar(&input, "input", defaultInput, usage)
  flag.StringVar(&input, "i", defaultInput, usage + " (shorthand)")

  flag.StringVar(&output, "output", defaultOutput, usage)
  flag.StringVar(&output, "o", defaultOutput, usage + " (shorthand)")

  flag.StringVar(&sample, "sample", defaultSample, usage)
  flag.StringVar(&sample, "s", defaultSample, usage + " (shorthand)")
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func main() {
    flag.Parse()

    input_file, err := os.Open(input)
    if err != nil { log.Fatal(err) }
    defer input_file.Close()

    output_file, err := os.Create(output)
    if err != nil { log.Fatal(err) }
    defer output_file.Close()

    reader := bufio.NewReader(input_file)
    writer := bufio.NewWriter(output_file)

    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF { break }
        writer.WriteString(string(line))
        if err != nil { log.Fatal(err) }
    }
}