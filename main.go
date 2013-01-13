package main

import (
  "log"
  "flag"
  "io"
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

    reader, input_file, _ := genReader(input)
    writer, output_file, _ := genWriter(output)
    defer input_file.Close()
    defer output_file.Close()

    for {
        line, _, err := reader.ReadLine()
        if err == io.EOF { break }

        _, err = writer.WriteString(string(line) + "\n")
        if err != nil { log.Fatal(err) }

        writer.Flush()
    }
}