package main

import (
	"bufio"
	"fmt"
	"os"
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

	/* 	var names []string */
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var ages2 map[string]int
	fmt.Println(ages2 == nil)
	fmt.Println(len(ages2) == 0)

	/* ages2["carol"] = 21 */

	if _, ok := ages["bob"]; !ok {
		fmt.Println("Key does not exist")
	}

	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
