package main

import (
    "strconv"
    "bufio"
    "os"
)

// exists returns whether the given file or directory exists or not
// Source: http://stackoverflow.com/a/10510783
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }

    return false, err
}

// genReader helper function acquires a bufio.Reader instance
// pointer to a O_RONLY mode openend file.
// You can use it to, for example, read a file line by line.
// It takes in parameter the path to the file you'd wanna
// acquire a Reader for, and returns a bufio.Reader instance pointer,
// a pointer to the opened os.File descriptor and error.
func GenReader(filepath string) (*bufio.Reader, *os.File, error) {
    input_file, err := os.Open(filepath)
    if err != nil { return nil, nil, err }

    reader := bufio.NewReaderSize(input_file, 8192)

    return reader, input_file, nil
}

// genWriter helper function acquires a bufio.Writer instance
// pointer to a write-mode openend file.
// You can use it to, for example, to write lines into a file.
// It takes in parameter the path to the file you'd wanna
// acquire a Reader for, and returns a bufio.Reader instance pointer,
// a pointer to the opened os.File descriptor and error.
// Plus, it opens the file in O_CREATE and O_WRONLY mode,
// With a 8192 bytes buffer.
func GenWriter(filepath string) (*bufio.Writer, *os.File, error) {
    output_file, err := os.OpenFile(filepath, os.O_CREATE | os.O_WRONLY, 0777)
    if err != nil { return nil, nil, err }

    writer := bufio.NewWriterSize(output_file, 8192)

    return writer, output_file, nil
}

// GenPartFilename helper function generates a part file
// name, adding extension, and number to the base file name
func GenPartFilename(basename string, ext string, num int) string{
    full_extension := "." + ext + strconv.Itoa(num)
    filename := basename + full_extension

    return filename
}