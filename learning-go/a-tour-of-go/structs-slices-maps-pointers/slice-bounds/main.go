package main

import (
	"fmt"
)

func main() {

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
	// s = s[low : high]
	// if you omit the low its default is 0
	// if you omit the high its value is the highest bound

	/*
	these slice expressions are all the same for this array var a [10]int:
	
	a[0:10]
	a[:10]
	a[0:]
	a[:]
	
	*/

}