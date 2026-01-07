package main

import (
	"fmt"
)

func main() {
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	stack = append(stack, 5)
	stack = append(stack, 6)
	stack = append(stack, 7)
	stack = append(stack, 8)
	stack = append(stack, 9)

	top := stack[len(stack)-1]   // top of stack
	stack = stack[:len(stack)-1] // pop, you are grabbing the slice from position 0 to len-1
	fmt.Println(stack, top)

	fmt.Println(remove(stack, 2))
	fmt.Println(removeNoOrder(stack, 2))
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeNoOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
