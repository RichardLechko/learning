package main

import (
	"fmt"
	"strconv"
)

func main() {
	// x := 123
	// y := fmt.Sprintf("%d", x)
	// fmt.Println(y, strconv.Itoa(x))
	// fmt.Println(y, strconv.FormatInt(int64(x), 10))
	// fmt.Println(y, strconv.FormatUint(uint64(x), 10))
	// fmt.Println(strconv.FormatInt(int64(x), 2))
	// fmt.Println(strconv.FormatUint(uint64(x), 2))
	x, _ := strconv.Atoi("123")
	y, _ := strconv.ParseInt("123", 10, 64)
	z, _ := strconv.ParseUint("123", 10, 64)
	fmt.Printf("x = %d\ty = %d\tz = %d\n", x, y, z)
}
