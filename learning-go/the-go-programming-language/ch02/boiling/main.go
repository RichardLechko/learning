// Boiling prints the boiling point of water
// this declares the package
package main

// this is an import statement
import (
	"fmt"
)

// this is a package level declaration, meaning it can be used in other .go files in the same directory
const boilingF = 212.0

func main() {

	// var "f" and "c" are local to function main. thus can ONLY be used in this function
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("Boiling point = %gF or %gC\n", f, c)
}