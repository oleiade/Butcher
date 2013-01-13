package main

import (
  "path"
  "bufio"
  "fmt"
  "log"
  "flag"
  "io"
  "os"
)

// /*
// ** Command line arguments
// */
var input string
var output string
var sample int

func init() {
  const (
    defaultInput = ""
    defaultOutput = ""
    defaultSample = 25
    usage = "butcher -input bigfatfile -output dir -sample linecountpersubfile"
  )

  flag.StringVar(&input, "input", defaultInput, usage)
  flag.StringVar(&input, "i", defaultInput, usage + " (shorthand)")

  flag.StringVar(&output, "output", defaultOutput, usage)
  flag.StringVar(&output, "o", defaultOutput, usage + " (shorthand)")

  flag.IntVar(&sample, "sample", defaultSample, usage)
  flag.IntVar(&sample, "s", defaultSample, usage + " (shorthand)")
}

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func main() {
    var writer                              *bufio.Writer
    var output_file                       *os.File
    var linenum, filenum               int = 0, 1

    // Parse command line arguments
    flag.Parse()

    // Acquire a Reader over input file
    reader, input_file, err := GenReader(input)
    defer input_file.Close()
    if err != nil {log.Fatal(err)}

    // If output dir does not exists yet, create it.
    if  val, err := exists(output) ; val == false {
      err = os.MkdirAll(output, 0777)
      if err != nil { log.Fatal(err) }
    }

    for {
        if (linenum % sample) == 0 {
            // Generate part file path
            partfile_name := GenPartFilename(input, "part", filenum)
            outfile_path := path.Join(output, partfile_name)
            fmt.Println(outfile_path)

            // Acquire a bufio.Writer pointing to part file
            writer, output_file, err = GenWriter(outfile_path)
            defer output_file.Close()
            if err != nil { log.Fatal(err) }

            filenum++
        }

        line, _, err := reader.ReadLine()
        if err == io.EOF { break }

        _, err = writer.WriteString(string(line) + "\n")
        if err != nil { log.Fatal(err) }

        writer.Flush()
        linenum++
    }
}