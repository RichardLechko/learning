package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

const debug = true

func main() {
	var w io.Writer
	fmt.Println(w)
	w = os.Stdout
	fmt.Println(w)
	w = new(bytes.Buffer)
	fmt.Println(w)
	w = nil

	var x any = time.Now()
	fmt.Println(x)

	var a io.Writer
	fmt.Printf("%T\n", a)

	a = os.Stdout
	fmt.Printf("%T\n", a)

	a = new(bytes.Buffer)
	fmt.Printf("%T\n", a)

	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ... use buf ...
	}
}

func f(out io.Writer) {
	// ... do something ...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
