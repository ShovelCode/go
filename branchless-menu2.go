package main

import (
	"fmt"
)

func option1() {
	fmt.Println("You selected Option 1.")
}

func option2() {
	fmt.Println("You selected Option 2.")
}

func option3() {
	fmt.Println("You selected Option 3.")
}

func option4() {
	fmt.Println("You selected Option 4.")
}

func option5() {
	fmt.Println("You selected Option 5.")
}

func option6() {
	fmt.Println("You selected Option 6.")
}

func invalidOption() {
	fmt.Println("Invalid Option.")
}

func main() {
	options := []func(){invalidOption, option1, option2, option3, option4, option5, option6}
	
	var userInput int

	fmt.Println("Select an option (1-6):")
	fmt.Scanln(&userInput)

	// Use the modulus operator to make sure userInput is a valid index
	// The "+ len(options)" ensures the userInput is not negative
	userInput = (userInput + len(options)) % len(options)

	options[userInput]()
}
