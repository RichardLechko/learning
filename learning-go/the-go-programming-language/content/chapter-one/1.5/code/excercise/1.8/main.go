// Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing. You might want to use strings.HasPrefix

package main

import (
	"fmt"
	"os"
	"strings"
	"io"
	"net/http"
)

func main() {

	s := "http://"
	if strings.HasPrefix(os.Args[1], s) == false {
		os.Args[1] = "http://" + os.Args[1]
	}
	
	for _, url := range os.Args[1:] {
	
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%v", b)
	}
}