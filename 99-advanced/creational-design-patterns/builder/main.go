package main

import "fmt"

/*
When to use Builder Pattern?
	- When object constructed is big and requires multiple steps
		- It helps in lessening the size of the constructor. Sometimes constructor is too verbose.
	- When different version of the same product needs to be created.
		- In this case: iglooBuilder, normalBuilder for igloo & house respectively
	- When half constructed final object should not exist. Either it is created 100% or not!

https://stackoverflow.com/questions/757743/what-is-the-difference-between-builder-design-pattern-and-factory-design-pattern

*/

func main() {
	normalBuilder := getBuilder("normal") // return normalBuilder struct
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder) // director ptr
	normalHouse := director.buildHouse()   // receiver method for director to use embedded builder to go through all build steps

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
