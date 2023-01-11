// Golang provides encapsulation at the package level
// Go doesn’t have any public,  private or protected keyword
//	Only mechanism to control the visibility is using the capitalized and non-capitalized formats
//	Non-capitalized identifiers are not exported
//	Capitalized Identifiers are exported
//		Structure
//		Structure's Method
//		Structure's Field
//		Function
//		Variable
//
//
//

package main

import "fmt"

var (
	CompanyName     = "test"     // Exported
	companyLocation = "somecity" // Not exported
)

type Person struct {
	Name string // Exported
	age  int    // Not exported
}

func (p *Person) GetAge() int { // Exported
	return p.age
}

func (p *Person) getName() string { // Not exported
	return p.Name
}

type company struct { // Not exported
}

func GetPerson() *Person { // Exported
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Println("Model Package:")
	fmt.Println(p.Name)
	fmt.Println(p.age)
	return p
}

func getCompanyName() string { // Not exported
	return CompanyName
}
