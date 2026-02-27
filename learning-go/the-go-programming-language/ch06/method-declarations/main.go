package main

import (
	"fmt"
	"gobook/ch06/method-declarations/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}
	perim := geometry.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println(perim.Distance())
}
