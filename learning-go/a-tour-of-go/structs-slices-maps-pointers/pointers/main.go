package main

import (
	"fmt"
)

func main() {

	i, j := 42, 2701
	
	p := &i // p holds the memory address of i

	fmt.Println(*p) // print the value that pointer p refers to
	*p = 21 // change the value in variable i's memory address
	fmt.Println(i)

	p = &j // p now holds the memory address of j
	*p = *p / 37 // the value that pointer p refers to will equal it's value divided by 37
	fmt.Println(j)
}