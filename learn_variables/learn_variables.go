package main

import "fmt"

func main() {
	// Go is a strongly typed language, meaning all variables
	// are bound to both a value and a type at compile time.

	// before you can use a variable in Go you must declare it
	var a int = 10 // long form of a variable declaration

	i := 10 // short hand for type inference, can't be used outside func

	var b byte = 10
	var c = 45
	// var i int -> produces i with the zero value
	// variables must be used (read) to avoid compilation error
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a)

	// long examples for declaration and initialization combinations
	var name, desc string = "Charles", "Developer"
	var radius int32 = 6378
	var mass float64 = 5.972e+24
	var IsActive bool = false
	var colors = []string{
		"yellow",
		"red",
		"blue",
		"green",
	}

	fmt.Println("Name:", name+" Desc:", desc)
	fmt.Println("Radius:", radius)
	fmt.Println("Mass:", mass)
	IsActive = true
	fmt.Println("Is Active:", IsActive)
	for _, color := range colors {
		fmt.Println("Color:", color)
	}

	// abbreviated examples for declaration and initialization combinations
	var FirstName, LastName = "Charles", "Hessifer"
	fmt.Println("Name: ", FirstName+" "+LastName)

	// numeric types
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// float32, float64 (not for currency)
	// bool -> true/false numeric 0 does not equate to false

	// Go forces numeric variables to be of same type when performing math
	j := 54
	k := 11

	fmt.Println(j + k)

	// example of type conversion
	var l int8 = 69
	var m float32 = 5.6

	fmt.Println(l + int8(m))

	// byte same as unit8
	// int same as int32 or int64
	// uint same as unit32 or uint64
}
