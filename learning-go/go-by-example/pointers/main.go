package main

import (
	"fmt"
)

/* pass by value */
/* this creates a copy of i and then changes the copy */
func zeroval(ival int) {
	ival = 0
}

/* pass by pointer */
/* this gets the address of i and then goes to the address and changes the value */
func zeroptr(iptr *int) {
	*iptr = 0
	/* "go to that address and change the value there" */
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)	

    fmt.Println("pointer:", &i)
}