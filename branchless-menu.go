package main

import "fmt"

// Define your menu options as functions
func option1() { fmt.Println("You selected option 1.") }
func option2() { fmt.Println("You selected option 2.") }
func option3() { fmt.Println("You selected option 3.") }
func option4() { fmt.Println("You selected option 4.") }
func option5() { fmt.Println("You selected option 5.") }
func option6() { fmt.Println("You selected option 6.") }
func defaultOption() { fmt.Println("Invalid option.") }

func main() {
	// Create an array of function pointers
	optionFuncs := []func(){defaultOption, option1, option2, option3, option4, option5, option6}

	// Simulate a user selection (for example, 3)
	// Make sure to validate user input in a real application
	userSelection := 3

	// Use array indexing for branchless execution
	optionFuncs[userSelection]()
}
