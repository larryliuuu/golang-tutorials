// Runtime Polymorphism in Go
// The same calculateTax is used in different contexts to calculate tax.
// When the compiler sees this call it delays which exact method to be called at run time.
// This magic happens behind the scene.

package main

import "fmt"

type taxSystem interface { // all structs that want to be typed as this must impl interface methods
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (i *usaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

func main() {
	indianTax := &indianTax{
		taxPercentage: 30,
		income:        1000,
	}
	singaporeTax := &singaporeTax{
		taxPercentage: 10,
		income:        2000,
	}
	usaTax := &usaTax{
		taxPercentage: 40,
		income:        500,
	}

	taxSystems := []taxSystem{indianTax, singaporeTax, usaTax}
	totalTax := calculateTotalTax(taxSystems)
	fmt.Printf("Total Tax is %d\n", totalTax)
}
func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax() // This is where runtime polymorphism happens!!!
	}
	return totalTax
}
