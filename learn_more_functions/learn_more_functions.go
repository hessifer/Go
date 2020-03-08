package main

import "fmt"

// does not conflict with addOne in main() - out of scope
func addOne(a int) int {
	return a + 1
}

func timesTwo(a int) int {
	return a * 2
}

// pass function as an argument to another function
func displayResults(a int, f func(int) int) {
	fmt.Println(f(a))
}

// return a function from a function
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

// funtion that takes in a function as an argument and returns a function
func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int { // closure
		b := f(a)
		return b * 2
	}
}

func main() {
	// assign function to variable, note no () since we are not invoking
	myAddOne := addOne
	fmt.Println(addOne(1))
	fmt.Println(myAddOne(1)) // treat variable like function

	// (anonymous function declaration) - no name specified after func keyword
	// in scope within main function
	addFive := func(a int) int {
		return timesTwo(a + 5) // invoke function inside of another function
	}
	fmt.Println(addOne(5))
	fmt.Println(addFive(4))

	// pass function as argument
	displayResults(2, addOne)
	displayResults(5, timesTwo)

	// closures is a function that has access to variables that exist in the
	// environments from where it was declared
	e := 9 // in the body of main
	fmt.Println("E before myAddTwenty: ", e)

	myAddTwenty := func(a int) {
		e = a * 20
	}
	myAddTwenty(40)
	fmt.Println("E after myAddTwenty: ", e)

	// return a function from a function
	myFuncA := makeAdder(1)
	myFuncB := makeAdder(2)

	fmt.Println(myFuncA(1))
	fmt.Println(myFuncB(1))

	// example of function passed for argument and returns a function
	doubleMyAddOne := makeDoubler(myAddOne)

	fmt.Println(doubleMyAddOne(1))
}
