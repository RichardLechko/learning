package main

import (
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    counts := make(map[string]int)
    
    for _, word := range words {
        counts[word]++
    }
    
    return counts
}

// Replace wc.Test with this
func Test(f func(string) map[string]int) {
    testCases := []string{
        "I am learning Go!",
        "The quick brown fox jumped over the lazy dog.",
        "I ate a donut. Then I ate another donut.",
    }
    
    for _, test := range testCases {
        fmt.Printf("Input: %q\n", test)
        result := f(test)
        fmt.Printf("Output: %v\n\n", result)
    }
}

func main() {
    Test(WordCount)
}