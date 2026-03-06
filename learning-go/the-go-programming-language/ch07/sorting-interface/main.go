package main

import (
	"fmt"
	"gobook/ch07/sorting-interface/sorting"
	"sort"
)

var tracks = sorting.Tracks

func main() {
	// names := []string{"John", "Tom", "Bob", "Michael", "Greg"}
	// fmt.Println(names)
	// sort.Sort(sort.StringSlice(names))
	// fmt.Println(names)
	sorting.PrintTracks(tracks)
	fmt.Println("---------------------------------------------------------------------")
	sort.Sort(sorting.ByArtist(tracks))
	sorting.PrintTracks(tracks)
	fmt.Println("---------------------------------------------------------------------")
	sort.Sort(sort.Reverse(sorting.ByArtist(tracks)))
	sorting.PrintTracks(tracks)
	fmt.Println("---------------------------------------------------------------------")
	sort.Sort(sorting.ByYear(tracks))
	sorting.PrintTracks(tracks)

	sort.Sort(sorting.CustomSort{T: tracks, CustLess: func(x, y *sorting.Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	sorting.PrintTracks(tracks)

	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))

}
