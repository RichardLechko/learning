package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2} // v = {1, 2}
	p := &v // p = the memory address of v
	p.X = 1e9 // change the value of pointer address pointing to X to 1e9
	fmt.Println(v)
}