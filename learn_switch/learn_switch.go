package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		word := os.Args[1]

		if word == "hello" {
			fmt.Println("Hi dude!")
		} else if word == "goodbye" {
			fmt.Println("See you when I see you!")
		} else if word == "greetings" {
			fmt.Println("Salutations!")
		} else {
			fmt.Println("I don't know what you said.")
		}

		fmt.Println()

		// if statements with more than 2 or 3 arms should become
		// switch statements
		switch word {
		case "hello":
			fmt.Println("Hi dude!")
		case "goodbye":
			fmt.Println("See you when I see you!")
		case "greetings":
			fmt.Println("Salutations!")
		default:
			fmt.Println("I don't know what you said!")
		}
	} else {
		fmt.Println("ERROR: expected two arguments go one.")
		os.Exit(-1)
	}

	w := "hello"

	/*
		Go switch statements do not fallthrough by default, you
		must explicitly use a "fallthrough" statement

		fallthrough indicates that control should flow from the end
		of the clause to the first statement of the next clause. Otherwise
		control flows to the end of the switch statement.

		fallthrough can appear in all clauses except for the last clause

	*/

	if len(os.Args) > 2 {
		switch l := len(os.Args[2]); {
		case w == "hello":
			fmt.Println("Greetings!")
			fallthrough
		case w == "dog":
			fmt.Println("Animal!")
		case l >= 5:
			fmt.Println("The word is greater than 5 bytes.")
		case l >= 1 && l < 5:
			fmt.Println("The word is less than 5 bytes.")
		default:
			fmt.Println("I just don't get what happened!")
		}
	}
}
