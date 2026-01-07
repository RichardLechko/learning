// Dup v2 prints the text of each line that appears more than once in the standard input, preceeded by its count

package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	// a map is a reference to the data structure made by make

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			// os.Open returns 2 values
			// 1. a open file (which is f *os.File) that will be used in countLines()
			// 2. a value of the built-in error type. if this second value is nil, the file was opened successfully. 

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				// if there is a error we print out the value of that error in natural format
				continue
			}
			countLines(f, counts)
			f.Close()
			// once we are done reading the file we close the file
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	// when we pass map into a function, the function receives a copy of the reference that we made with make
	// so any changes made in this function will be visible in the callers map reference, which is our main func
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}