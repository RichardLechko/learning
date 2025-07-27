package main

import (
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	//				 0, 1, 2, 3, 4, 5

	var s []int = primes[1:4]
	// s -> primes[low : high] -> primes[1 : 4] -> primes[1] + primes[2] + primes[3]
	fmt.Println(s)
}