package main

import (
	"fmt"
	"gobook/ch06/method-declarations/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))

	var origin geometry.Point

	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy
	scaleP(2)
	fmt.Println(p)
	scaleP(3)
	fmt.Println(p)
	scaleP(10)
	fmt.Println(p)

}
