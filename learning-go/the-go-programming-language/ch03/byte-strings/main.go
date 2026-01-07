package main

import (
	"fmt"
)

func main() {
	s := "abc"
	b := []byte(s)
	b[0] = byte('z')
	s2 := string(b)
	fmt.Println(s2)
}