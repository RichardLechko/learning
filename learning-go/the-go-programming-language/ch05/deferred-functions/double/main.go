package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("double(%d) = %d\n", 4, double(4))
	_ = doubleDefer(4)
	fmt.Println(triple(4))
}

func double(x int) int {
	return x + x
}

func doubleDefer(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func failedDefer(filenames []string) (err error) {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		// ... process f ...
	}
	return nil
}

func goodDefer(filenames []string) (err error) {
	for _, filename := range filenames {
		err := doFile(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ... process f ...
	return nil
}
