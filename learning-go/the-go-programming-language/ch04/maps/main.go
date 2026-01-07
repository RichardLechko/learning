package main

import (
	"fmt"
	"sort"
)

func main() {

	// ages := make(map[string]int) // this maps strings to integers
	// ages := map[string]int{}

	// ages := map[string]int{
	// 	"alice":   31,
	// 	"charlie": 34,
	// }

	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 34

	ages := make(map[string]int)
	ages["alice"] = 32
	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages)

	ages["bob"] = ages["bob"]
	fmt.Println(ages["bob"])
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages["bob"])

	ages["bob"] += 1
	fmt.Println(ages["bob"])
	ages["bob"]++
	fmt.Println(ages["bob"])

	ages["alice"] = 32
	ages["charlie"] = 34
	ages["tom"] = 60
	ages["john"] = 44

	// _ = &ages["bob"] // compile error: cannot take address of map element

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	fmt.Println("\n\n\n")

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

}
