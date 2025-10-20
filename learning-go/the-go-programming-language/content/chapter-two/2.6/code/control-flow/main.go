package main

import "fmt"

func main() {
	if 1 == 3 {
		x := 1
		y := 2
		fmt.Println(x, y)
		if 1 != 3 {
			a := 3
			b := 4
			fmt.Println(x, y, a, b)
		}
	}
	
}
