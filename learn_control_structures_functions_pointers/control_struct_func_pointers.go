package main

import "fmt"

func main() {
	a := 10

	// if / else example
	if a > 5 {
		fmt.Println("a is bigger than 5")
	} else {
		fmt.Println("a is less than or equal to 5")
	}

	// for loop example
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// while condition true
	j := 0

	for j < 10 {
		fmt.Println(j)
		j = j + 1
	}
	fmt.Println()

	// while True (invinite loop)
	k := 0

	for {
		fmt.Println(k)
		k = k + 1
		if k > 4 {
			break
		}
	}
	fmt.Println()

	// for range loop
	s := "miles"

	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
	fmt.Println()
	name := "miles"
	revName(name)
}

func revName(name string) {
	t := name
	tLen := len(t)

	for i := tLen - 1; i >= 0; i-- {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}
