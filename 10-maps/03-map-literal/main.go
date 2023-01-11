package main

import "fmt"

func main() {
	var m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, // Comma is necessary
	}
	var m2 = map[string]int{}
	if m2 != nil {
		fmt.Println("m2 is not nil")
	}

	fmt.Println(m2)

	fmt.Println(m)
}
