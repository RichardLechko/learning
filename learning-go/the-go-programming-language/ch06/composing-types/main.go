package main

import (
	"fmt"
	"gobook/ch06/composing-types/coloredpoint"
	"image/color"
)

func main() {
	var cp coloredpoint.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = coloredpoint.ColoredPoint{
		Point: coloredpoint.Point{X: 1, Y: 1},
		Color: red,
	}
	var q = coloredpoint.ColoredPoint{
		Point: coloredpoint.Point{X: 5, Y: 4},
		Color: blue,
	}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}
