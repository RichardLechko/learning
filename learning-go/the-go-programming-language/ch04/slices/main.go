package main

import (
	"fmt"
)

func main() {

	// this is an array that has all months 1-12
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	// we then can create a slice of the underlying array and grab all months in Q2
	Q2 := months[4:7]

	// then we grab all months in summer
	summer := months[6:9]

	fmt.Println(Q2)
	fmt.Println(summer)
	// this will print:
	// [April May June]
	// [June July August]
	// You can also see that we grabbed an overlapping element, June, and got no errors

	// This will print out any months that are in Q2 and in Summer, which just prints out "June appears in both"
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// This will result in an error because we are trying to slice outside of the range, this will throw an error: panic.
	// fmt.Println(summer[:20])

	// This will result in extending the slice to include more months
	fmt.Println(summer[:5]) // [June July August September October]

	// You can see that even though we only passed in a slice, it still changed the underlying array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // [5 4 3 2 1 0]

	s := []int{0, 1, 2, 3, 4, 5}
	// This rotates 's' left by two positions
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

	var z []int    // len is 0 and z is nil
	z = nil        // len is 0 and z is nil
	z = []int(nil) // len is 0 and z is nil
	z = []int{}    // len is 0 and z is not nil
}

// this function reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {

	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

/*

Slices are variable-length sequences and all elements have the same type.

They are written like []T

Slice is a data structure that gives access to a subsequence (or all) of the elements of an array, there are three components that go into a slice:
- pointer, which points to the first element
- length, which is # of slice elements
- capacity, which is the # of elements from the start of the slice to the end of the underlying array

Slices contain a pointer to an element of an array, meaning if you pass a slice to a function, the function can modify the slice which in turn modifys the original array.

Here is the difference between initializing a slice and an array:
- Array -> a := [...]int{0, 1, 2, 3, 4, 5}
- Slice -> s := []int{0, 1, 2, 3, 4, 5}
The main difference is that you do not need to provide a size to create a slice since it's a dynamic data structure

Slices are not comparable, so you can not use something like == to compare two slices.
- If you have a []byte slice, you can use `bytes.Equal`
- If it's not a []byte slice, you need to do the comparison yourself, which we did in `func equal`
- the only legal slice comparison is to compare it against nil to see if the slice is empty or not, but you should not do that. use len == 0 to see if a slice is empty

"The zero value of a slice type is nil. A nil slice has no underlying array. The nil slice has length and capacity zero."


*/
