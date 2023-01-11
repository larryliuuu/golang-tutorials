/* Compile Time Polymorphism in Go

* Go does not support Method Overloading
* Go does not support Function Overloading
* Go does not support Operator Overloading

Variadic Function provides a way to achieve the same result w/ increase code complexity
*/

package main

import "fmt"

type maths struct{}

/*
*  The two add methods below are not allowed as one overloads the other
 */

// func (m *maths) add(a, b int) int {
// 	return a + b
// }

// func (m *maths) add(a, b, c int) int {
// 	return a + b + c
// }

func (m *maths) add(numbers ...int) int { // Variadic Function
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func main() {
	m := &maths{}

	fmt.Printf("Result: %d\n", m.add(2, 3))
	fmt.Printf("Result: %d\n", m.add(2, 3, 4))
}
