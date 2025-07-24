package main

import (
	"fmt"
	"os"
)

func main() {

	// fmt.Println(strings.Join(os.Args[1:], " "))
	// this is a good way for formatting

	fmt.Println(os.Args[1:])
	// this is the most optimal if we do not care about formatting
}