// take temperature as celsius from user and
// display it in fahrenheit
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var celsiusTemp string

	fmt.Print("Enter the temperature in Celsius: ")
	for scanner.Scan() {
		celsiusTemp = scanner.Text()
		break
	}

	if cTemp, err := strconv.ParseFloat(celsiusTemp, 64); err == nil {
		fmt.Printf("%.2f degrees celsius is %.2f degrees fahrenheit.\n", cTemp, celsiusToFahrenheit(float64(cTemp)))
	} else {
		fmt.Printf("celToFahr: %v\n", err)
	}
}

// conver celsius to fahrenheit and return result
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}
