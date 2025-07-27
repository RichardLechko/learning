package main

import (
	"fmt"
)

type Vertext struct {
	X, Y int
}

var (
	v1 = Vertext{1, 2}
	v2 = Vertext{X: 1}
	v3 = Vertext{} // int uninitialized are 0 -> {0 0}
	p = &Vertext{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}