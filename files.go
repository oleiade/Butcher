package main

import (
    "os"
    "bufio"
)

func genReader(filepath string) (*bufio.Reader, *os.File, error) {
    input_file, err := os.Open(filepath)
    if err != nil { return nil, nil, err }

    reader := bufio.NewReaderSize(input_file, 8192)

    return reader, input_file, nil
}

func genWriter(filepath string) (*bufio.Writer, *os.File, error) {
    output_file, err := os.OpenFile(filepath, os.O_CREATE | os.O_WRONLY, 0777)
    if err != nil { return nil, nil, err }

    writer := bufio.NewWriterSize(output_file, 8192)

    return writer, output_file, nil
}