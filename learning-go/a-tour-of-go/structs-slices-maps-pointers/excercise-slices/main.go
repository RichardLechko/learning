package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {

	picture := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		picture[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {

			picture[y][x] = uint8(x ^ y)
		}
	}

	return picture

}

func displayPicture(pic [][]uint8) {
    fmt.Println("Picture pattern:")
    for _, row := range pic {
        for _, pixel := range row {
            fmt.Printf("%3d ", pixel)
        }
        fmt.Println()
    }
}

func main() {
    // Test with a small picture
    result := Pic(8, 8)
    displayPicture(result)
    
    // Print structure info
    fmt.Printf("\nStructure: %d rows, each with %d columns\n", len(result), len(result[0]))
    fmt.Printf("Type: %T\n", result)
}