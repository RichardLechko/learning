package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	// w = time.Second // compile error, time.Duration does not implement io.Writer

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// rwc = new(bytes.Buffer) // compile error, bytes.Buffer does not implement io.ReadWriteCloser
	fmt.Printf("%v & %v\n", w, rwc)

	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
}
