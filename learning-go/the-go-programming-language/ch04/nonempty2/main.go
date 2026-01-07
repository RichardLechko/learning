package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data))
	fmt.Printf("%q\n", data)
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of underlying array
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
