package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	chapterOne()
}

func chapterOne() {
	echoProgramInefficient()
	fmt.Println()
	echoProgram()
	fmt.Println()
	exerciseOneDotOne()
	fmt.Println()
	exerciseOneDotTwo()
	fmt.Println()
	exerciseOneDotThree()
	fmt.Println()
}

func echoProgram() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echoProgramInefficient() {
	/* Example of Quadratic Process

	In this context quadratic process means that the cost of the process is proportional
	to the square of the input size. This is because strings are concatenated using +=
	operator which is expensive in Go as strings are immutable and a new string must
	be created in memory every time you concatenate. More efficient ways to concatenate
	strings include writing to bytes.Buffer and converting it to string, or using
	strings.Join function

	*/

	start := time.Now()
	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	microSecs := time.Since(start).Microseconds()
	fmt.Printf("It took %d microSecs to execute echoProgramInefficient().", microSecs)
}

func exerciseOneDotOne() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func exerciseOneDotTwo() {
	for i, v := range os.Args[1:] {
		fmt.Println(i+1, v)
	}
}

func exerciseOneDotThree() {
	start := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))
	microSecs := time.Since(start).Microseconds()
	fmt.Printf("It took %d microSecs to execute exerciseOneDotThree().", microSecs)
}
