package main

// struct embedding is similar to inheritance

import (
	"fmt"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

// we pass base into the container struct, thus container gets all the attributes of base.

func main() {

	co := container {
		base: base {
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// we can access co.num even though the num field only exists in base, but container inherited it.

	fmt.Println("also num:", co.base.num)
	// you can access the num more explicitly

	fmt.Println("describe:", co.describe())
	// container also inherits all of base's methods

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}