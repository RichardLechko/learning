package main

import (
	"fmt"
)

func main() {
	fmt.Println(comma("12345"))
	fmt.Println(comma("999999999"))
	fmt.Println(comma("82311230812"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}