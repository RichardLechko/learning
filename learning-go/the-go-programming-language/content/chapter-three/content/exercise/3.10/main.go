package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234"))
}

func comma(s string) string {
	var buf bytes.Buffer
	var counter = 0
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("i = %d\tv = %c\n", i, s[i])
		if counter != 0 {
			if counter%3 == 0 {
				buf.WriteString(",")
			}
			fmt.Fprintf(&buf, "%d", s[i])
		}

		counter++
	}
	return buf.String()
	/* return "" */
}

/*

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

*/
