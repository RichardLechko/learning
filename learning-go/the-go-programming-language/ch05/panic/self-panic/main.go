package main

import (
	"fmt"
)

func main() {
	switch s := suit(drawCard()); s {
	case "Spades":
		// ....
	case "Hearts":
		// ...
	case "Diamonds":
		// ...
	case "Clubs":
		// ...
	default:
		panic(fmt.Sprintf("invalid suit %q:, s")) // Joker?
	}
}
