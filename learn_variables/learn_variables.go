package main

import "fmt"

func main() {
	// var i int = 10
	i := 10 // short hand for type inference, can't be used outside func

	var b byte = 10
	var c = 45
	// var i int -> produces i with the zero value
	// variables must be used (read) to avoid compilation error

	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(c)

	// numeric types
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// float32, float64 (not for currency)
	// bool -> true/false numeric 0 does not equate to false

	// Go forces numeric varialbes to be of same type when performing math
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
