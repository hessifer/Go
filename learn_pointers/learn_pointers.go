/*
Go Is Not C
- Go has built-in array and string types
- Go has GC
- go is strongly typed
*/

package main

import "fmt"

// pass pointer to function as argument
func setTo10(a *int) {
	*a = 10
}

func main() {
	// NOTE: Go is a pass by value language

	a := 10
	b := &a // referrence to 'a'
	c := a
	fmt.Println(a, b, *b, c) // '*' dereference to get value

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30 // dereference to set value
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	// declare variable to pointer_type
	// var g *int // default to zero value of 'nil' (abscense of value, produces a panic)

	g := new(int) // new makes a pointer for type passed in and 'allocates' memory

	fmt.Println(g)  // print memory location
	fmt.Println(*g) // derefernce value in memory location

	// pass pointer to function as argument
	setTo10(g)
	fmt.Println(g)  // print memory location
	fmt.Println(*g) // derefernce value in memory location (we have set value from nil to 10)

	myFavNumber := 5
	setTo10(&myFavNumber)
	fmt.Println(myFavNumber)

}
