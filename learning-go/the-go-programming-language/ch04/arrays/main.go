package main

import (
	"fmt"
)

func main() {
	// because we did not initialize this array with values, the array will use the zero-value, which is 0 for numbers
	// this array is then: [0, 0, 0]
	var a [3]int
	/* You can access the element in an array using subscript notation. */
	fmt.Println(a[0])        // this prints the first element
	fmt.Println(a[len(a)-1]) // this prints the last element which is a[2]

	// this will print out the indices and the values:
	// 0 0, 1 0, 2 0
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// we use a blank identifier since in this case we only want to print out the elements, not the indices
	// 0 0 0
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// if you want to create array literals to initialize an array with values
	//var q [3]int = [3]int{1, 2, 3} // this array will be [1, 2, 3]
	var r [3]int = [3]int{1, 2} // this array will be [1, 2, 0]
	fmt.Println(r[2])           // this prints the last element which is 0

	// the array literal above is really verbose, you can use short-hand declaration and an ellipsis:
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // %T prints out the type of the value

	// q := [3]int{1, 2, 3}
	// q = [4]int{1, 2, 3, 4} this will throw an error bc you can not assign [4]int to [3]int

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	// You can also create an array by using the "iota" constant generator and create a key value pairing
	symbol := [...]string{USD: "Dollar", EUR: "Euro", GBP: "Pound", RMB: "Yuan"}
	fmt.Println(RMB, symbol[RMB]) // this will print "3 Yuan"

	// this will create an array where the frist 99 elements are 0 and the 100th element is -1
	z := [...]int{99: -1}
	fmt.Println(z)

	// if the elements in an array can be comparable that means the array as a whole can be compared
	// since integers are comparable, that means an array of integers is also comparable
	aa := [2]int{1, 2}
	bb := [...]int{1, 2}
	cc := [2]int{1, 3}
	fmt.Println(aa == bb, aa == cc, bb == cc) // this will print "true false false"
	// dd := [3]int{1, 2}
	// fmt.Println(aa == dd) this will throw a compiler error because you are trying to compare a [2]int to a [3]int which are different types

}

// We pass in the pointer of this byte array as an argument, meaning the function is not receiving a copy of the array it's actually changing the original.
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

/*

An array is a fixed length with 0 or more elements that are all the same type.

The size of an array must be a constant expression, meaning the compiler must know the value at compile time

the size of an array is a part of the type, thus [1, 2, 3, 4] and [1, 2, 3] are both arrays holding integers but they are both different types

In Go, passing arguments into a function creates a copy. So the function receives a copy, meaning passing in large arrays is inefficient
So, if you want to pass in an array and have the function change the original array and not a copy, pass in the pointer instead as an argument

"Using a pointer to an array is efficient and allows the called function to mutate the caller's variable".

*/
