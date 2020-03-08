package main

import (
	"fmt"
)

/* Go Functions
1. input parameters separated by comma
2. parameter name followed by type
3. return type written prior to open \{
4. location of function definition does not matter
5. no default values for parameters
6. no optional parameters
7. no named parameters
8. no function overloading
9. by default call by value, meaning a copy of the variable's
   value is passed to function (original value is unmodified)
10. functions are first class citizens in go, and can be passed
   around similar to
*/

// function that takes two int arguments and returns a
// single int value
func sumTwoNumbers(num1 int, num2 int) int {
	return num1 + num2
}

// functionthat takes two int arguments and returns
// 'two' int values
func productOf(num1 int, num2 int) (int, int) {
	return num1 * num2, num1 % num2
}

func main() {
	number1 := 5
	number2 := 10
	total := sumTwoNumbers(number1, number2)

	fmt.Printf("Sum of %d + %d is %d\n", number1, number2, total)

	product, remainder := productOf(number1, number2)
	fmt.Printf("Product of %d and %d is %d, the remainder is %d\n", number1, number2, product, remainder)

	// to ignore a return value use '_'
	_, answer := productOf(number1, number2)
	fmt.Println(answer)
}
