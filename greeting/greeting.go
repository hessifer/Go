// ask user for name and displays a simple greeting
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		name := scanner.Text()
		if name != "" {
			fmt.Println("Hello,", name)
		} else {
			fmt.Println("Hi")
		}
		break
	}
}
