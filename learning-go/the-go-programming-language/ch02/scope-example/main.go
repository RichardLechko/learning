package main

import "fmt"

func main() {

	var a = 3
	fmt.Printf("Outside for loop: %d\n", a)
	for i := range 5 {
		var a = 1
		fmt.Println(a, i)
	}
}