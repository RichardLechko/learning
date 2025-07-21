package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for {

		newZ := z - (z*z - x) / (2*z)

		if math.Abs(newZ - z) < 0.0000001 {
			break
		}
		fmt.Println(z)
		z = newZ
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}