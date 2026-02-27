package main

import (
	"fmt"
	"gobook/ch06/method-declarations/geometry"
	"net/url"
)

func main() {
	perim := geometry.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	fmt.Println(perim.Distance())

	r := &geometry.Point{X: 1, Y: 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := geometry.Point{X: 1, Y: 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)

	q := geometry.Point{X: 1, Y: 2}
	(&q).ScaleBy(2)
	fmt.Println(q)

	t := geometry.Point{X: 1, Y: 2}
	t.ScaleBy(2)
	fmt.Println(t)

	geometry.Point{X: 1, Y: 2}.Distance(q) // Point
	pptr.ScaleBy(2)                        // *Point

	p.ScaleBy(2) // implicit (&p)

	pptr.Distance(q) // implicit (*pptr)

	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3")
}
