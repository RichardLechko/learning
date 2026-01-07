package main

import(
	"fmt"
	"os"
	"strconv"
	"conversions/lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		mm := lengthconversions.Millimeter(t)
		cm := lengthconversions.Centimeter(t)
		m  := lengthconversions.Meter(t)
		km := lengthconversions.Kilometer(t)
		fmt.Println("KILOMETER CONVERSION")
		fmt.Println("=================================================")
		fmt.Printf("KM: %s \t M: %s\n", km, lengthconversions.KMToM(km))
		fmt.Printf("KM: %s \t CM: %s\n", km, lengthconversions.KMToCM(km))
		fmt.Printf("KM: %s \t MM: %s\n", km, lengthconversions.KMToMM(km))
		fmt.Printf("=================================================\n\n")
		fmt.Println("METER CONVERSION")
		fmt.Println("=================================================")
		fmt.Printf("M: %s \t KM: %s\n", m, lengthconversions.MToKM(m))
		fmt.Printf("M: %s \t CM: %s\n", m, lengthconversions.MToCM(m))
		fmt.Printf("M: %s \t MM: %s\n", m, lengthconversions.MToMM(m))
		fmt.Printf("=================================================\n\n")
		fmt.Println("CENTIMETER CONVERSION")
		fmt.Println("=================================================")
		fmt.Printf("CM: %s \t KM: %s\n", cm, lengthconversions.CMToKM(cm))
		fmt.Printf("CM: %s \t M: %s\n", cm, lengthconversions.CMToM(cm))
		fmt.Printf("CM: %s \t MM: %s\n", cm, lengthconversions.CMToMM(cm))
		fmt.Printf("=================================================\n\n")
		fmt.Println("MILLIMETER CONVERSION")
		fmt.Println("=================================================")
		fmt.Printf("MM: %s \t KM: %s\n", mm, lengthconversions.MMToKM(mm))
		fmt.Printf("MM: %s \t M: %s\n", mm, lengthconversions.MMToM(mm))
		fmt.Printf("MM: %s \t CM: %s\n", mm, lengthconversions.MMToCM(mm))
		fmt.Printf("=================================================\n\n")
	}
}