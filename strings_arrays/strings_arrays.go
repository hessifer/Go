package main

import "fmt"

func main() {
	// var s string // 0 value assigned to s variable

	s := "Hello World!" // assign string literal to s

	fmt.Println(s)

	// Go treats a string as an immutable sequence of bytes
	firstName := "Charles"
	lastName := "Hessifer"
	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	// Get the first letter byte representation of firstName and display to the screen
	fmt.Println(firstName[0])

	// Convert the above byte to a string
	fmt.Println(string(firstName[0]))

	// Display last half of firstName
	firstNameLength := len(firstName)
	fmt.Println("First Name Length: ", firstNameLength)

	midPoint := int(firstNameLength / 2)
	fmt.Println("Mid Point: ", midPoint)

	fmt.Println(string(firstName[midPoint:]))

	// Display last characther in lastName
	fmt.Println(string(lastName[len(lastName)-1]))

	// NOTE: Go stores characters in UTF-8 consuming 1 - 4 bytes
	// len() will return the bytes of all charachters combined, keep in mind UTF-8 and some may be
	//     larger than 1 byte

	// Go has a special numeric type called a Rune which is equivalent to an i32 (signed value)
	var r rune = 127757
	fmt.Println(string(r))

	t := 'c'

	// i32 value
	fmt.Println("Rune: ", t)

	// casted to string
	fmt.Println("Rune: ", string(t))

	// Array in Go is similar to other languages
	// - fixed size
	// - contains same data type
	// - sequece of values
	var nums [3]int
	nums[0] = 2
	nums[1] = 5
	nums[2] = 9

	fmt.Println(nums[0], nums[1], nums[2])

	// NOTE: In Go the length of the array is part of the type
	// var nums2 [5]int = nums   --> this would not compile since we have a variable of type [3]int and [5]int
	// no type conversion
	// if you wanted to have a function in Go to sum all the values in an array you would need a function to handle
	// the permutation of type []int, this is not ideal which is why the 'slice' data type is used more often

}

// WTF is a rune, really a char type.....
