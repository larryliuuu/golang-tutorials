// implementing type hierarchy using interfaces and struct
//  * Use embedding both on struct and on interface level
//  * Go does not provide method overriding.

package main

import "fmt"

type iAnimal interface {
	breathe()
}
type animal struct {
}

func (a *animal) breathe() {
	fmt.Println("Animal breathe")
}

type iAquatic interface {
	iAnimal // embedded interface within interface
	swim()
}
type aquatic struct {
	animal // embedded struct within struct
}

func (a *aquatic) swim() {
	fmt.Println("Aquatic swim")
}

type iNonAquatic interface {
	iAnimal // embedded interface within interface
	walk()
}
type nonAquatic struct {
	animal // embedded struct within struct
}

func (a *nonAquatic) walk() {
	fmt.Println("Non-Aquatic walk")
}

type shark struct {
	aquatic
}
type lion struct {
	nonAquatic
}

func main() {
	shark := &shark{}
	checkAquatic(shark)
	checkAnimal(shark)
	lion := &lion{}
	checkNonAquatic(lion)
	checkAnimal(lion)

	shark.breathe()
	shark.swim()
	lion.breathe()
	lion.walk()
}
func checkAquatic(a iAquatic)       {}
func checkNonAquatic(a iNonAquatic) {}
func checkAnimal(a iAnimal)         {}
