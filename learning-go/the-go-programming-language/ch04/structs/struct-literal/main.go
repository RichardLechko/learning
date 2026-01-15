package main

import "fmt"

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

func main() {
	pp := &Point{X: 10, Y: 20}
	fmt.Println(ScaleX(pp, 10))

	p := Point{X: 1, Y: 2}
	q := Point{X: 2, Y: 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func ScaleX(p *Point, factor int) int {
	return p.X * factor
}
