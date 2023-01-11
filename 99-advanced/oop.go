// Go prefers composition to inheritance: allows embedding of struct into other struct
// Go does not support type inheritance

// https://golangbyexample.com/oop-inheritance-golang-complete/

//
// Combine usage of struct and interface
//		* interface "iBase" has "say" method
//		* "check" accepts an argument of type "iBase"
//		* "clear" is a property which is of type function in the base struct
//			* This makes it possible for child's "clear" to be invoked, and note that of base's
package main

import "fmt"

type iBase interface {
	say()
}
type base struct {
	color string
	clear func()
}

func (b *base) say() {
	b.clear()
}

type child struct {
	base  // embedding struct within struct
	style string
}

func check(b iBase) {
	b.say()
}
func main() {
	base := base{color: "Red",
		clear: func() {
			fmt.Println("Clear from child's function")
		}}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
}
