// get advanced generics examples from youtube video: https://www.youtube.com/watch?v=Kf5LLi7Ti9A&ab_channel=CodeWithRyan
// type and its receiver from youtube video: https://www.youtube.com/watch?v=Kf5LLi7Ti9A&ab_channel=CodeWithRyan

// go mod vendor

// slices
/*
A slice will create a new array and copy the contents of the old one into the new one only as a result of append operation,
and only if the capacity of the underlying array is not sufficient to hold the new elements.
Your operations never require a change in the underlying array capacity, so the slices all point to the same underlying array.
*/

// Pointer to struct
// Go lets you directly access the fields of a struct through the pointer without explicit dereference.

/*
Struct - exported vs. unexported

Any struct type that starts with a capital letter is exported and accessible from outside packages. Similarly, any struct field that starts with a capital letter is exported.
On the contrary, all the names starting with a small letter are visible only inside the same package.
*/

/*
Structs are value types

Structs are value types. When you assign one struct variable to another, a new copy of the struct is created and assigned.
Similarly, when you pass a struct to another function, the function gets its own copy of the struct.
*/

/*
METHODS vs. FUNCTIONS

Methods w/ Pointer receivers - why we have these

With a value receiver, the method operates on a copy of the value passed to it. Therefore, any modifications done to the receiver inside the method is not visible to the caller.
Ex: func (gp GeometricProgression) NthTerm(n int) float64

Go also allows you to define a method with a pointer receiver. Inside the below example, you can modify gp struct instance's fields/properties
Ex: func (gp *GeometricProgression) NthTerm(n int) float64


*Methods with Pointer receivers vs Functions with Pointer arguments*
A method with a pointer receiver can accept both a pointer and a value as the receiver argument.
But, a function with a pointer argument can only accept a pointer

*Methods w/ value receivers vs. Functions w/ value arguments*
A method with a value receiver can accept both a value and a pointer as the receiver argument.
But, a function with a value argument can only accept a value

(Pretty much always use a pointer receiver instead of value receiver) - you'll most likely always want to change what is underneath....

*Where can methods be defined*
For you to be able to define a method on a receiver, the receiver type must be defined in the same package.
Go doesn’t allow you to define a method on a receiver type that is defined in some other package (this includes built-in types such as int as well).

*Methods on non struct types*
Go allows you to define methods on non-struct types too. In the following example, I’ve defined a method called reverse() on the type MyString.
*/

/* Interfaces

Under the hood, an interface value can be thought of as a tuple consisting of a value and a concrete type:
(value, type)

*/

/*

By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.

recv Channels are blocking
- can synchronize channels this way via making a done channel, then having groutine do `done<-true` and have main thread do `<-done`

channel directions: When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
This specificity increases the type-safety of the program. Note the argument syntax.

func ping(pings chan<- string, msg string) ... pings only sends

func pong(pings <-chan string, pongs chan<- string) ... pings only receives, pong only sends


Go’s `select` lets you wait on multiple channel operations.

*/

/*
Timers and Tickers
*/

/*
Mutexes: https://gobyexample.com/mutexes

Stateful Goroutines instead of mutexes: https://gobyexample.com/stateful-goroutines

*/

// alias fmt.Println: var p = fmt.Println

/*


 */

package golang_tutorials
