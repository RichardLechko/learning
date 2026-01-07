package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "Dollar", EUR: "Euro", GBP: "Pound", RMB: "Yuan"}
	fmt.Println(RMB, symbol[RMB])

	r := [...]int{99: -1}
	fmt.Println(r)
}
